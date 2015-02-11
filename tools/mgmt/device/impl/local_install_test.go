package impl_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"v.io/core/veyron2/naming"
	"v.io/core/veyron2/security"
	"v.io/core/veyron2/services/mgmt/application"
	"v.io/core/veyron2/services/mgmt/device"

	"v.io/core/veyron/tools/mgmt/device/impl"
)

func createFile(t *testing.T, path string, contents string) {
	if err := ioutil.WriteFile(path, []byte(contents), 0700); err != nil {
		t.Fatalf("Failed to create %v: %v", path, err)
	}
}

func TestInstallLocalCommand(t *testing.T) {
	shutdown := initTest()
	defer shutdown()

	tape := NewTape()
	server, endpoint, err := startServer(t, gctx, tape)
	if err != nil {
		return
	}
	defer stopServer(t, server)
	// Setup the command-line.
	cmd := impl.Root()
	var stdout, stderr bytes.Buffer
	cmd.Init(nil, &stdout, &stderr)
	deviceName := naming.JoinAddressName(endpoint.String(), "")
	const appTitle = "Appo di tutti Appi"
	binary := os.Args[0]
	fi, err := os.Stat(binary)
	if err != nil {
		t.Fatalf("Failed to stat %v: %v", binary, err)
	}
	binarySize := fi.Size()
	for i, c := range []struct {
		args         []string
		stderrSubstr string
	}{
		{
			[]string{deviceName}, "incorrect number of arguments",
		},
		{
			[]string{deviceName, appTitle}, "missing binary",
		},
		{
			[]string{deviceName, appTitle, "a=b"}, "missing binary",
		},
		{
			[]string{deviceName, appTitle, "foo"}, "binary foo not found",
		},
		{
			[]string{deviceName, appTitle, binary, "PACKAGES", "foo"}, "foo not found",
		},
	} {
		c.args = append([]string{"install-local"}, c.args...)
		if err := cmd.Execute(c.args); err == nil {
			t.Fatalf("test case %d: wrongly failed to receive a non-nil error.", i)
		} else {
			fmt.Fprintln(&stderr, "ERROR:", err)
			if want, got := c.stderrSubstr, stderr.String(); !strings.Contains(got, want) {
				t.Errorf("test case %d: %q not found in stderr: %q", i, want, got)
			}
		}
		if got, expected := len(tape.Play()), 0; got != expected {
			t.Errorf("test case %d: invalid call sequence. Got %v, want %v", got, expected)
		}
		tape.Rewind()
		stdout.Reset()
		stderr.Reset()
	}
	emptySig := security.Signature{Purpose: []uint8{}, Hash: "", R: []uint8{}, S: []uint8{}}
	emptyBlessings := security.WireBlessings{}
	cfg := device.Config{"someflag": "somevalue"}

	testPackagesDir, err := ioutil.TempDir("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testPackagesDir)
	pkgFile1 := filepath.Join(testPackagesDir, "file1.txt")
	createFile(t, pkgFile1, "1234567")
	pkgFile2 := filepath.Join(testPackagesDir, "file2")
	createFile(t, pkgFile2, string([]byte{0x01, 0x02, 0x03, 0x04}))
	pkgDir1 := filepath.Join(testPackagesDir, "dir1")
	if err := os.Mkdir(pkgDir1, 0700); err != nil {
		t.Fatalf("Failed to create dir1: %v", err)
	}
	createFile(t, filepath.Join(pkgDir1, "f1"), "123")
	createFile(t, filepath.Join(pkgDir1, "f2"), "456")
	createFile(t, filepath.Join(pkgDir1, "f3"), "7890")

	pkgFile3 := filepath.Join(testPackagesDir, "file3")
	createFile(t, pkgFile3, "12345")
	pkgFile4 := filepath.Join(testPackagesDir, "file4")
	createFile(t, pkgFile4, "123")
	pkgDir2 := filepath.Join(testPackagesDir, "dir2")
	if err := os.Mkdir(pkgDir2, 0700); err != nil {
		t.Fatalf("Failed to create dir2: %v", err)
	}
	createFile(t, filepath.Join(pkgDir2, "f1"), "123456")
	createFile(t, filepath.Join(pkgDir2, "f2"), "78")
	pkg := application.Packages{
		"overridepkg1": application.PackageSpec{File: pkgFile3},
		"overridepkg2": application.PackageSpec{File: pkgFile4},
		"overridepkg3": application.PackageSpec{File: pkgDir2},
	}

	for i, c := range []struct {
		args         []string
		config       device.Config
		packages     application.Packages
		expectedTape interface{}
	}{
		{
			[]string{deviceName, appTitle, binary},
			nil,
			nil,
			InstallStimulus{
				"Install",
				appNameAfterFetch,
				nil,
				nil,
				application.Envelope{
					Title:     appTitle,
					Binary:    binaryNameAfterFetch,
					Signature: emptySig,
					Publisher: emptyBlessings,
				},
				map[string]int64{"binary": binarySize}},
		},
		{
			[]string{deviceName, appTitle, binary},
			cfg,
			nil,
			InstallStimulus{
				"Install",
				appNameAfterFetch,
				cfg,
				nil,
				application.Envelope{
					Title:     appTitle,
					Binary:    binaryNameAfterFetch,
					Signature: emptySig,
					Publisher: emptyBlessings,
				},
				map[string]int64{"binary": binarySize}},
		},
		{
			[]string{deviceName, appTitle, "ENV1=V1", "ENV2=V2", binary, "FLAG1=V1", "FLAG2=V2"},
			nil,
			nil,
			InstallStimulus{
				"Install",
				appNameAfterFetch,
				nil,
				nil,
				application.Envelope{
					Title:     appTitle,
					Binary:    binaryNameAfterFetch,
					Signature: emptySig,
					Publisher: emptyBlessings,
					Env:       []string{"ENV1=V1", "ENV2=V2"},
					Args:      []string{"FLAG1=V1", "FLAG2=V2"},
				},
				map[string]int64{"binary": binarySize}},
		},
		{
			[]string{deviceName, appTitle, "ENV=V", binary, "FLAG=V", "PACKAGES", pkgFile1, pkgFile2, pkgDir1},
			nil,
			pkg,
			InstallStimulus{"Install",
				appNameAfterFetch,
				nil,
				nil,
				application.Envelope{
					Title:     appTitle,
					Binary:    binaryNameAfterFetch,
					Signature: emptySig,
					Publisher: emptyBlessings,
					Env:       []string{"ENV=V"},
					Args:      []string{"FLAG=V"},
				},
				map[string]int64{
					"binary":                        binarySize,
					"packages/file1.txt":            7,
					"packages/file2":                4,
					"packages/dir1":                 10,
					"overridepackages/overridepkg1": 5,
					"overridepackages/overridepkg2": 3,
					"overridepackages/overridepkg3": 8,
				},
			},
		},
	} {
		const appId = "myBestAppID"
		tape.SetResponses([]interface{}{InstallResponse{appId, nil}})
		if c.config != nil {
			jsonConfig, err := json.Marshal(c.config)
			if err != nil {
				t.Fatalf("test case %d: Marshal(%v) failed: %v", i, c.config, err)
			}
			c.args = append([]string{fmt.Sprintf("--config=%s", string(jsonConfig))}, c.args...)
		}
		if c.packages != nil {
			jsonPackages, err := json.Marshal(c.packages)
			if err != nil {
				t.Fatalf("test case %d: Marshal(%v) failed: %v", i, c.packages, err)
			}
			c.args = append([]string{fmt.Sprintf("--packages=%s", string(jsonPackages))}, c.args...)
		}
		c.args = append([]string{"install-local"}, c.args...)
		if err := cmd.Execute(c.args); err != nil {
			t.Fatalf("test case %d: %v", i, err)
		}
		if expected, got := fmt.Sprintf("Successfully installed: %q", naming.Join(deviceName, appId)), strings.TrimSpace(stdout.String()); got != expected {
			t.Fatalf("test case %d: Unexpected output from Install. Got %q, expected %q", i, got, expected)
		}
		if got, expected := tape.Play(), []interface{}{c.expectedTape}; !reflect.DeepEqual(expected, got) {
			t.Errorf("test case %d: Invalid call sequence. Got %#v, want %#v", i, got, expected)
		}
		tape.Rewind()
		stdout.Reset()
		stderr.Reset()
	}
}

package godirs

import (
  "bufio"
  "bytes"
  "io"
  "fmt"
  "path/filepath"
  "os"
  "os/exec"
  "strings"
  "syscall"
)

func exists(path string) bool {
  _, err := os.Stat(path)
  return !os.IsNotExist(err)
}

func readLines(path string) ([]string, error) {
  f, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer f.Close()

  lines := []string{}
  r := bufio.NewReader(f)
  var b bytes.Buffer
  for {
    l, p, err := r.ReadLine()
    if err == io.EOF {
      break
    } else if err != nil {
      return nil, err
    }

    b.Write(l)
    if !p {
      lines = append(lines, b.String())
      b.Reset()
    }
  }

  return lines, nil
}

func projectPaths() ([]string, error) {
  cwd, err := os.Getwd()
  if err != nil {
    return nil, err
  }

  for d := cwd; d != "/"; d = filepath.Dir(d) {
    cfg := filepath.Join(d, ".gaan")
    if !exists(cfg) {
      continue
    }

    p, err := readLines(cfg)
    if err != nil {
      return nil, err
    }

    for i := 0; i < len(p); i++ {
      p[i] = filepath.Join(d, p[i])
    }

    return p, nil
  }

  return []string{}, nil
}

func FindGoRoot() (string, error) {
  if p, err := exec.LookPath("go"); err == nil {
    return filepath.Dir(filepath.Dir(p)), nil
  }

  p := "/usr/local/go"
  _, err := os.Stat(p)
  if err == nil {
    return p, nil
  }

  return "", err
}

func Run(goroot, tool string) (int, error) {
  if p, err := projectPaths(); err == nil {
    os.Setenv("GOPATH", strings.Join(p, ":"))
  }

  args := os.Args
  cmd := make([]string, len(args))
  cmd[0] = fmt.Sprintf("%s/bin/%s", goroot, tool)
  for i, n := 1, len(args); i < n; i++ {
    cmd[i] = args[i]
  }

  proc, err := os.StartProcess(cmd[0],
    cmd,
    &os.ProcAttr{
      "",
      os.Environ(),
      []*os.File{os.Stdin, os.Stdout, os.Stderr},
      nil})
  if err != nil {
    return 0, err
  }

  s, err := proc.Wait()
  if err != nil {
    return 0, err
  }

  return s.Sys().(syscall.WaitStatus).ExitStatus(), nil
}

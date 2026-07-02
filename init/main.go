package main
import "os/exec"
func main() { exec.Command("open", "-a", "Calculator").Run() }

package judge

import (
	"fmt"
	"testing"
)

func TestCompileCPP(t *testing.T) {
	code := "#include <iostream>\n\nint main()\n{\n  std::cout << \"hello world\" << std::endl;\n  return 0;\n}"
	fmt.Println(compileCPP(1, code, "/home/cu1/judgeEnvironment/submissions"))
}

func TestCompileC(t *testing.T) {
	code := "#include <stdio.h>\n\nint main()\n{\n  printf(\"hello world\\n\");\n  return 0;\n}"
	fmt.Println(compileC(2, code, "/home/cu1/judgeEnvironment/submissions"))
}

func TestCompilePython2(t *testing.T) {
	code := "print(\"hello world\")"
	fmt.Println(compilePython2(3, code, "/home/cu1/judgeEnvironment/submissions"))
}

func TestCompilePython3(t *testing.T) {
	code := "print(\"hello world\")"
	fmt.Println(compilePython3(4, code, "/home/cu1/judgeEnvironment/submissions"))
}

func TestCompileGo(t *testing.T) {
	code := "package main\n\nimport (\n        \"fmt\"\n)\n\nfunc main() {\n   fmt.Println(\"hello world\")\n}"
	fmt.Println(compileGo(5, code, "/home/cu1/judgeEnvironment/submissions"))
}

func TestCompileJava(t *testing.T) {
	code := "import java.util.*;\npublic class Main{\n\tpublic static void main(String args[]){\n\t\tScanner cin = new Scanner(System.in);\n\t\tint a, b;\n\t\t\n\t\t\tSystem.out.println(\"hello world\");\n\t\t\n\t}\n}"
	fmt.Println(compileJava(6, code, "/home/cu1/judgeEnvironment/submissions"))
}

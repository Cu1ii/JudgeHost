package xoj

import (
	"fmt"
	"testing"
)

func TestCompileCPP(t *testing.T) {
	code := "#include <ostrea>\n\nint main()\n{\n  std::cout << \"hello world\" << std::endl;\n  return 0;\n}"
	fmt.Println(compileCPP(1, code, "/home/cu1/XOJ/submission", "1"))
}

func TestCompileC(t *testing.T) {
	code := "#include <stdio.h>\n\nint main()\n{\n  printf(\"hello world\\n\");\n  return 0;\n}"
	fmt.Println(compileC(2, code, "/home/cu1/XOJ/submission", "2"))
}

func TestCompilePython2(t *testing.T) {
	code := "print(\"1\")"
	fmt.Println(compilePython2(4, code, "/home/cu1/XOJ/submission", "3"))
}

func TestCompilePython3(t *testing.T) {
	code := "print(\"4\")"
	fmt.Println(compilePython3(4, code, "/home/cu1/XOJ/submission", "4"))
}

func TestCompileGo(t *testing.T) {
	code := "package main\n\nimport (\n        \"fmt\"\n)\n\nfunc main() {\n   fmt.Println(\"hello world\")\n}"
	fmt.Println(compileGo(5, code, "/home/cu1/XOJ/submission", "5"))
}

func TestCompileJava(t *testing.T) {
	code := "import java.util.*;\npublic class Main{\n\tpublic static void main(String args[]){\n\t\tScanner cin = new Scanner(System.in);\n\t\tint a, b;\n\t\twhile (cin.hasNext()){\n\t\t\ta = cin.nextInt(); b = cin.nextInt();\n\t\t\tSystem.out.println(a + b);\n\t\t}\n\t}\n}"
	fmt.Println(compileJava(6, code, "/home/cu1/XOJ/submission", "6"))
}

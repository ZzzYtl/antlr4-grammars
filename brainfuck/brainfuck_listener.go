// Code generated from brainfuck.g4 by ANTLR 4.7.2. DO NOT EDIT.

package brainfuck // brainfuck
import "github.com/antlr/antlr4/runtime/Go/antlr"

// brainfuckListener is a complete listener for a parse tree produced by brainfuckParser.
type brainfuckListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterOpcode is called when entering the opcode production.
	EnterOpcode(c *OpcodeContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitOpcode is called when exiting the opcode production.
	ExitOpcode(c *OpcodeContext)
}

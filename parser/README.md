# Parser

See if we can build an almost single pass compiler for Go doing syntax
directed translation, without any complex data structure (no syntax
tree), only lists of tokens.

The goal is to have the shortest and simplest path from source to
bytecode.

## Design

The input of parser is a list of tokens produced by the scanner.
Multiple tokens are processed at once. The minimal set to get
meaningful results (not an error or nil) is a complete statement.

The output of parser is also a list of tokens, to be consumed by
the compiler to produce bytecode. The output tokens set is identical
to the bytecode instructions set except that:

- code locations may be provided as labels instead of numerical
  values,
- memory locations for constants and variables may be provided as
  symbol names instead of numerical values.

## Status

Go language support:

- [x] named functions
- [x] anonymous functions (closures)
- [ ] methods
- [x] internal function calls
- [x] external function calls (calling runtime symbols in interpreter)
- [ ] export to runtime
- [ ] builtin calls (new, make, copy, delete, len, cap, ...)
- [ ] out of order declarations
- [x] arbirtrary precision constants
- [x] basic types
- [ ] complete numeric types
- [x] function types
- [ ] variadic functions
- [ ] pointers
- [x] structures
- [ ] embedded structures
- [ ] recursive structures
- [ ] literal composite objects
- [ ] interfaces
- [x] arrays, slices
- [ ] maps
- [ ] deterministic maps
- [ ] channel types
- [ ] channel operations
- [x] var defined by assign :=
- [x] var assign =
- [x] var declaration
- [x] type declaration
- [x] func declaration
- [x] const declaration
- [x] iota expression
- [ ] defer statement
- [ ] recover statement
- [ ] go statement
- [x] if statement (including else and else if)
- [x] for statement
- [x] switch statement
- [ ] type switch statement
- [x] break statement
- [x] continue statement
- [ ] fallthrough statement
- [x] goto statement
- [x] label statement
- [ ] select statement
- [x] binary operators
- [x] unary operators
- [x] logical operators && and ||
- [ ] assign operators
- [x] operator precedence rules
- [x] parenthesis expressions
- [x] call expressions
- [x] index expressions
- [x] selector expressions
- [ ] slice expressions
- [ ] type convertions
- [ ] type assertions
- [ ] parametric types (generic)
- [ ] type parametric functions (generic)
- [ ] type constraints (generic)
- [ ] type checking
- [ ] comment pragmas
- [ ] package import
- [ ] modules

Other items:

- [x] REPL
- [x] multiline statements in REPL
- [ ] completion, history in REPL
- [x] eval strings
- [x] eval files (including stdin, ...)
- [x] debug traces for scanner, parser, compiler, bytecode vm
- [x] simple interpreter tests to exercise from source to execution
- [ ] compile time error detection and diagnosis
- [ ] stack dump
- [ ] symbol tables, data tables, code binded to source lines
- [ ] interactive debugger: breaks, continue, instrospection, ...
- [ ] machine level debugger
- [ ] source level debugger
- [ ] replay debugger, backward instruction execution
- [ ] vm monitor: live memory / code display during run
- [ ] stdlib wrappers a la yaegi
- [ ] system and environment sandboxing
- [ ] build constraints (arch, sys, etc)
- [ ] test command (running go test / benchmark / example files)
- [ ] skipping / including test files
- [ ] test coverage
- [ ] fuzzy tests for scanner, vm, ...

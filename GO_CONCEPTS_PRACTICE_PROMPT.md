# Go Concepts Practice Prompt (from this repository)

Copy-paste this into GitHub Copilot Chat whenever you want targeted Go practice based on the concepts covered in this repository.

---

## Prompt to paste in Copilot Chat

I practiced Go before and I am revising now.  
Use the concept list below to generate practice coding tasks and solutions in Go.

### What I want from you
1. Ask me which concept(s) I want to practice first.
2. Generate **3 levels** of practice for the selected concept:
   - Level 1: Beginner
   - Level 2: Intermediate
   - Level 3: Slightly challenging
3. For each level, provide:
   - Problem statement
   - Input/output examples
   - Hints (without full solution first)
4. Wait for my attempt.
5. When I ask, provide:
   - Clean idiomatic Go solution
   - Time/space complexity
   - Common mistakes and edge cases
6. After each solved problem, suggest 1 follow-up variation.
7. Keep problems practical and small enough to finish in 15–30 minutes each.

### Concepts to use (from my repository)

1. Program structure and setup (`package main`, `main()`, imports, modules)
2. Variables, declarations, and zero values
3. Constants (typed and untyped)
4. Basic/predeclared types (int, float, bool, string, byte, rune, complex)
5. Explicit type conversion
6. Strings and Unicode (UTF-8 bytes vs runes, string iteration)
7. Control flow (`if`, `for`, `switch`, `defer`)
8. Functions (parameters, returns, multiple returns, named returns)
9. Error handling (`error`, `errors.New`, branching on errors)
10. Scope rules (package/function/block scope, shadowing, blank identifier `_`)
11. Arrays (declaration, indexing, value semantics, comparison)
12. Slices (slicing, append, len/cap, copy, nil vs empty, preallocation)
13. Maps (CRUD operations, existence checks, iteration)
14. Structs and composition/embedding
15. Pointers (`&`, `*`, pointer mutation, pointers to structs)
16. Methods (value vs pointer receivers)
17. Interfaces (definition and implementation via method sets)
18. Concurrency basics (goroutines, channels, `sync.WaitGroup`, `sync.Mutex`)
19. Time/performance basics (`time.Now`, `time.Since`, `time.Sleep`)
20. Standard library practice (`fmt`, `math`, `strings`, `sync`, `time`, `net/http`, etc.)
21. HTTP basics (`http.HandleFunc`, handlers, server startup)

### Practice mode rules
- Start from easy and progressively increase difficulty.
- Include tests for at least intermediate problems.
- Prefer idiomatic Go style.
- Mention when a concept overlaps with another (e.g., slices + pointers, goroutines + channels).
- If I say **“quiz mode”**, do not give solution until I explicitly ask.

---

## Optional quick-start requests

You can also paste one of these in Copilot Chat:

- “Give me 5 mixed Go practice tasks from this concept list, increasing difficulty.”
- “Test me only on slices, maps, and structs today.”
- “Give me concurrency-focused tasks using goroutines, channels, WaitGroup, and Mutex.”
- “Revision sprint: 30-minute set covering variables, functions, control flow, and slices.”
- “Interview mode: ask one question at a time, evaluate my answer, then continue.”


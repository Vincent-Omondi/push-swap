# Push-Swap Project File Structure

## Main Structure

- **main.go**: Entry point for the project. Initializes the CLI interface and delegates tasks to either the `pushswap` or `checker` packages based on command-line arguments.

## pushswap Package

- **pushswap.go**: Contains the main logic for the `push-swap` program. Handles input argument parsing, validation, and generates the minimal set of sorting instructions. Utilizes functions from `sort.go` and `validate.go`.
  
- **sort.go**: Implements the non-comparative sorting algorithm required by the project specifications.
  
- **validate.go**: Provides functions to validate input arguments for the `pushswap` program, ensuring integer validity, uniqueness, and other error conditions as specified.

## checker Package

- **checker.go**: Manages the `checker` program's logic. Reads stack inputs and instructions from standard input, executes the instructions, and verifies if the stack is sorted correctly using functions from `execute.go` and `validate.go`.
  
- **execute.go**: Implements functions to execute stack manipulation instructions (`sa`, `sb`, `pa`, `pb`, etc.) for the `checker` program.
  
- **validate.go**: Provides validation functions tailored for the `checker` program, ensuring stack state correctness and handling input errors.

## tests Directory

- **pushswap_test.go**: Contains unit tests for the `pushswap` package. Tests cover edge cases for sorting algorithms and input validation.
  
- **checker_test.go**: Holds unit tests for the `checker` package. Tests ensure correct handling and verification of stack operations and instruction executions.

## Additional Considerations

- **Error Handling**: Both `pushswap` and `checker` programs handle errors gracefully, printing error messages to standard error when necessary.
  
- **Separation of Concerns**: Each file and package has a distinct responsibility, promoting readability and maintainability of the codebase.
  
- **Command Line Interface**: `main.go` facilitates command-line argument parsing and task delegation to appropriate packages (`pushswap` or `checker`).

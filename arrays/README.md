# Arrays Problems

This directory contains various array-based algorithm problems, each organized in its own folder with individual test cases.

## Structure

Each problem has its own folder containing:
- `problem_name.go` - Problem implementation with multiple approaches
- `main.go` - Test cases specific to that problem

## Available Problems

- `single_number/` - Find the single number in an array
- `sort_colors/` - Sort an array of colors (Dutch National Flag problem)
- `rearrange_elements_by_sign/` - Rearrange array with alternating positive/negative elements
- `majority_element/` - Find the majority element in an array
- `maximum_consecutive_ones/` - Find maximum consecutive 1's in a binary array
- `maximum_subarray_sum/` - Find maximum sum of contiguous subarray (Kadane's algorithm)
- `next_permutation/` - Find the next lexicographically greater permutation
- `best_time_to_buy_and_sell_stock/` - Find maximum profit from stock trading
- `find_missing_element/` - Find missing number in array
- `longest_subarray_with_sum_k/` - Find longest subarray with given sum
- `two_sum/` - Find two numbers that add up to target
- `general_array_problems/` - General array manipulation utilities

## Using the Makefile

### Basic Commands

```bash
# Show help and available commands
make help

# Run all problems
make all

# Run a quick test of key problems
make quick

# Validate all problems compile
make validate

# Clean build artifacts
make clean

# Show problem statistics
make stats
```

### Individual Problem Commands

```bash
# Run specific problems
make single_number
make sort_colors
make rearrange_elements
make majority_element
make max_consecutive_ones
make max_subarray_sum
make next_permutation
make best_time_to_buy
make find_missing_element
make longest_subarray
make two_sum
make general_problems
```

### Generic Run Command

```bash
# Run any problem by name
make run PROBLEM=single_number
make run PROBLEM=sort_colors
```

## Bash Completion

For enhanced productivity, you can enable bash completion:

```bash
# Load completion (temporary)
source completion.sh

# Make completion permanent
echo 'source /path/to/dsa/arrays/completion.sh' >> ~/.bashrc
```

With completion enabled, you can:
- Type `make <TAB>` to see all available targets
- Type `make single_<TAB>` to auto-complete problem names
- Type `make run PROBLEM=<TAB>` to see available problems

## Running Individual Problems

You can also run problems directly:

```bash
# Navigate to problem folder
cd single_number

# Run the problem
go run .

# Or build and run
go build .
./single_number
```

## Development

### Adding New Problems

1. Create a new folder: `mkdir new_problem`
2. Add implementation: `new_problem.go`
3. Add tests: `main.go`
4. Update Makefile with new target
5. Update completion script if needed

### Testing

Each problem folder contains comprehensive test cases that verify:
- Multiple algorithm approaches (brute force, optimized, etc.)
- Edge cases and boundary conditions
- Performance characteristics

## Examples

```bash
# Quick development workflow
make validate          # Check all problems compile
make quick            # Run key problems
make single_number    # Test specific problem
make clean            # Clean up
```

## Performance

The Makefile provides performance insights:
- Individual test execution times
- Overall test suite statistics
- Compilation validation

Each problem includes multiple algorithmic approaches with time/space complexity analysis in the code comments.

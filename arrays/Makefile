# Arrays Problems Makefile
# Usage: make <problem_name> or make help

.PHONY: help all clean test

# Default target
help:
	@echo "Available commands:"
	@echo "  make help                    - Show this help message"
	@echo "  make all                     - Run all problems"
	@echo "  make clean                   - Clean build artifacts"
	@echo "  make test                    - Run all tests"
	@echo ""
	@echo "Individual problems:"
	@echo "  make single_number           - Run single number problem"
	@echo "  make sort_colors             - Run sort colors problem"
	@echo "  make rearrange_elements      - Run rearrange elements by sign problem"
	@echo "  make majority_element       - Run majority element problem"
	@echo "  make max_consecutive_ones    - Run maximum consecutive ones problem"
	@echo "  make max_subarray_sum       - Run maximum subarray sum problem"
	@echo "  make next_permutation       - Run next permutation problem"
	@echo "  make best_time_to_buy       - Run best time to buy and sell stock problem"
	@echo "  make find_missing_element   - Run find missing element problem"
	@echo "  make longest_subarray       - Run longest subarray with sum k problem"
	@echo "  make two_sum                - Run two sum problem"
	@echo "  make general_problems       - Run general array problems"
	@echo ""
	@echo "Quick test commands:"
	@echo "  make quick                  - Run a few key problems"
	@echo "  make validate               - Validate all problems compile"

# Run all problems
all: single_number sort_colors rearrange_elements majority_element max_consecutive_ones max_subarray_sum next_permutation best_time_to_buy find_missing_element longest_subarray two_sum general_problems

# Quick test - run a few key problems
quick: single_number sort_colors two_sum max_subarray_sum

# Validate all problems compile
validate:
	@echo "Validating all problems compile..."
	@for dir in single_number sort_colors rearrange_elements_by_sign majority_element maximum_consecutive_ones maximum_subarray_sum next_permutation best_time_to_buy_and_sell_stock find_missing_element longest_subarray_with_sum_k two_sum general_array_problems; do \
		echo "Checking $$dir..."; \
		cd $$dir && go build . && cd ..; \
	done
	@echo "All problems compile successfully! ✅"

# Individual problem targets
single_number:
	@echo "Running Single Number problem..."
	@cd single_number && go run .

sort_colors:
	@echo "Running Sort Colors problem..."
	@cd sort_colors && go run .

rearrange_elements:
	@echo "Running Rearrange Elements by Sign problem..."
	@cd rearrange_elements_by_sign && go run .

majority_element:
	@echo "Running Majority Element problem..."
	@cd majority_element && go run .

max_consecutive_ones:
	@echo "Running Maximum Consecutive Ones problem..."
	@cd maximum_consecutive_ones && go run .

max_subarray_sum:
	@echo "Running Maximum Subarray Sum problem..."
	@cd maximum_subarray_sum && go run .

next_permutation:
	@echo "Running Next Permutation problem..."
	@cd next_permutation && go run .

best_time_to_buy:
	@echo "Running Best Time to Buy and Sell Stock problem..."
	@cd best_time_to_buy_and_sell_stock && go run .

find_missing_element:
	@echo "Running Find Missing Element problem..."
	@cd find_missing_element && go run .

longest_subarray:
	@echo "Running Longest Subarray with Sum K problem..."
	@cd longest_subarray_with_sum_k && go run .

two_sum:
	@echo "Running Two Sum problem..."
	@cd two_sum && go run .

general_problems:
	@echo "Running General Array Problems..."
	@cd general_array_problems && go run .

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@find . -name "*.exe" -delete
	@find . -name "*.test" -delete
	@echo "Clean complete! ✅"

# Test all problems (alias for 'all')
test: all

# Show available problems
list:
	@echo "Available problems:"
	@ls -1d */ | sed 's|/||' | sort

# Run a specific problem with error handling
run:
	@if [ -z "$(PROBLEM)" ]; then \
		echo "Usage: make run PROBLEM=<problem_name>"; \
		echo "Available problems:"; \
		ls -1d */ | sed 's|/||' | sort; \
		exit 1; \
	fi
	@if [ ! -d "$(PROBLEM)" ]; then \
		echo "Problem '$(PROBLEM)' not found!"; \
		echo "Available problems:"; \
		ls -1d */ | sed 's|/||' | sort; \
		exit 1; \
	fi
	@echo "Running $(PROBLEM) problem..."
	@cd $(PROBLEM) && go run .

# Development helpers
dev-setup:
	@echo "Setting up development environment..."
	@go mod tidy
	@echo "Development setup complete! ✅"

# Show problem statistics
stats:
	@echo "Problem Statistics:"
	@echo "Total problems: $$(ls -1d */ | wc -l)"
	@echo "Problems with tests: $$(find . -name "main.go" | wc -l)"
	@echo ""
	@echo "Problem breakdown:"
	@for dir in */; do \
		problem=$$(echo $$dir | sed 's|/||'); \
		test_count=$$(find $$dir -name "*.go" | wc -l); \
		echo "  $$problem: $$test_count Go files"; \
	done

#!/bin/bash

# Bash completion script for Arrays Makefile
# Usage: source completion.sh

_makefile_completion() {
    local cur prev opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    
    # Available make targets
    opts="help all clean test quick validate list stats dev-setup single_number sort_colors rearrange_elements majority_element max_consecutive_ones max_subarray_sum next_permutation best_time_to_buy find_missing_element longest_subarray two_sum general_problems run"
    
    # If the previous word is 'run', complete with problem names
    if [[ ${prev} == "run" ]]; then
        local problems=$(ls -1d */ 2>/dev/null | sed 's|/||' | tr '\n' ' ')
        COMPREPLY=( $(compgen -W "${problems}" -- ${cur}) )
        return 0
    fi
    
    # If the previous word is 'make' and current word starts with a dash, complete with options
    if [[ ${prev} == "make" && ${cur} == -* ]]; then
        COMPREPLY=( $(compgen -W "-j -k -n -s -t -v -w -C -f -o -W" -- ${cur}) )
        return 0
    fi
    
    # Complete with available targets
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
}

# Register the completion function
complete -F _makefile_completion make

# Also provide completion for the run command specifically
_make_run_completion() {
    local cur prev
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    
    if [[ ${prev} == "PROBLEM=" ]]; then
        local problems=$(ls -1d */ 2>/dev/null | sed 's|/||' | tr '\n' ' ')
        COMPREPLY=( $(compgen -W "${problems}" -- ${cur}) )
        return 0
    fi
}

echo "Bash completion for Arrays Makefile loaded!"
echo "Usage examples:"
echo "  make <TAB>                    # Show all available targets"
echo "  make run PROBLEM=<TAB>        # Show available problems"
echo "  make single_<TAB>             # Auto-complete problem names"
echo ""
echo "To make this permanent, add 'source $(pwd)/completion.sh' to your ~/.bashrc or ~/.zshrc"

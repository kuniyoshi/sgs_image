# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

### Build and Run
```bash
# Run the main application
make run
# or
go run ./cmd/app

# Build the application
make build
# or
go build -o bin/app ./cmd/app

# Run sandbox experiments
make sandbox
# or
go run ./cmd/sandbox
```

## Architecture Overview

This is a scenario scripting DSL experiment that separates scenario logic from game engine implementation.

### Module Structure
The project uses Go workspaces with three modules:
- `scenario/` - Core scenario library (independent)
- `cmd/app/` - Game engine implementation
- `cmd/sandbox/` - Experimental testing ground

### Key Design Principles

1. **Scenario Independence**: The scenario module must remain completely independent from any game engine implementation. It defines what happens, not how it's rendered.

2. **Transition-Based State**: Transitions represent end states, not animations. This ensures:
   - Idempotency (can be applied multiple times safely)
   - Resumability from any point
   - Engine-agnostic state representation

3. **Animation Responsibility**: The game engine (cmd/app) handles:
   - Animation interpolation between transitions
   - User input processing
   - State management and rendering

### Core Types in scenario/
- `Vector3`: 3D coordinates (X, Y, Z)
- `Camera`: Position and Direction vectors
- `Transition`: State changes (currently camera movements)

### Application Flow (cmd/app)
1. Initialize scenario with `scenario.Begin()`
2. Get transitions with `scenario.Progress()`
3. Animate between transitions (10 steps per transition)
4. Handle user input: Enter to progress, 'q' to skip
5. Check completion with `scenario.IsEnd()`

### Development Notes
- The README is in Japanese and explains the experimental nature
- Focus is on proving scenario/engine separation concept
- Same scenario should work with different game engines
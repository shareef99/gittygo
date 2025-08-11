# GittyGo

├──gittygo/
│
├── cmd/
│ └── gittygo/
│ └── main.go # Entry point: initializes app + main event loop
│
├── internal/
│ ├── ui/ # All UI code (Gio layouts, drawing, widgets)
│ │ ├── app.go # Main UI struct (root layout, navigation)
│ │ ├── components/ # Reusable UI elements
│ │ │ ├── button.go
│ │ │ ├── sidebar.go
│ │ │ └── top-bar.go
│ │ ├── pages/ # Each screen/page in the app
│ │ │ ├── home.go
│ │ │ ├── repo.go
│ │ │ └── commits.go
│ │ └── theme.go # Colors, typography, spacing
│ │
│ ├── git/ # Git-related logic
│ │ ├── repo.go # Repo opening/closing
│ │ ├── commits.go # Commit history fetching
│ │ ├── branches.go # Branch operations
│ │ └── status.go # Status/staging logic
│ │
│ ├── state/ # Global state + events
│ │ ├── app_state.go # Holds current app state (selected repo, branch, etc.)
│ │ └── actions.go # State update functions
│ │
│ └── util/ # Helpers and shared utilities
│ ├── file.go
│ └── format.go
│
├── go.mod
└── go.sum

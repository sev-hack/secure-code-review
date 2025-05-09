# Path Traversal

Although the code tries to clean the path via `filepath.Clean` method, it's still vulnerable because this method requires the path starting with a slash, otherwise it won't be cleaned.

Simply replace `cleanPath := filepath.Clean(filename)` with `cleanPath := filepath.Clean("/" + filename)` in order to fix the problem.

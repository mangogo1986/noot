# Git Commit Message Specification

## Background and Importance

During development, each Git commit requires a Commit message. A standardized Commit message provides several benefits:

- **Easy browsing and tracking**: Quickly locate and trace work history.
- **Automated Changelog generation**: Automatically generate version difference notes during releases.

Current issues include:

- **Inconsistent styles**: Everyone uses different styles, making it hard to browse.
- **Missing information**: Some commits lack descriptions, making it difficult to understand their purpose.

Standardizing Commit messages resolves these issues without adding complexity or learning costs.

---

## Commit Message Format

We use a simple, easy-to-follow format, supporting both English and Chinese:

```plaintext
<type>(<scope>): <subject>
```

- **type** (required): The type of change, detailed in the table below.
- **scope** (optional): The scope of the change, such as the data layer, view layer, or directory name.
- **subject** (required): A brief description of the change.

### Commit Types

| Type       | Description                             |
|------------|-----------------------------------------|
| `feat`     | New feature                             |
| `fix`      | Bug fix                                 |
| `docs`     | Documentation changes                  |
| `style`    | Code style changes (non-functional)    |
| `refactor` | Code refactoring (no new features/bugs)|
| `perf`     | Performance optimization               |
| `test`     | Adding tests                           |
| `chore`    | Build or tooling changes               |
| `revert`   | Revert changes                         |
| `build`    | Packaging                              |

**Example:**

```plaintext
feat(miniprogram): add support for mini-program template messages
```

---

## Integration into Projects

### 1. Install Dependencies

```bash
npm i commitlint --save-dev
npm i @commitlint/config-conventional --save-dev
npm i husky --save-dev
npm install commitizen --save-dev
commitizen init cz-customizable --save --save-exact
```

### 2. Configure Commitlint

Create a `commitlint.config.js` file in the project root with the following content:

```javascript
module.exports = {
  extends: ['@commitlint/config-conventional'],
  rules: {
    'type-enum': [2, 'always', [
      "feat", "fix", "docs", "style", "refactor",
      "perf", "test", "chore", "revert", "build"
    ]],
    'subject-case': [0] // Disable subject-case validation
  },
};
```

### 3. Configure Husky

Add the following configuration to `package.json`:

```json
"husky": {
  "hooks": {
    "commit-msg": "commitlint -E HUSKY_GIT_PARAMS"
  }
}
```

---

### 4. Configure Commitizen

Add the following configuration to `package.json`:

```json
"config": {
  "commitizen": {
    "path": "./node_modules/cz-customizable"
  }
}
```

Create a `.cz-config.js` file in the project root with the following content:

```javascript
'use strict';

module.exports = {
  types: [
    {value: 'feat', name: 'feat: New feature'},
    {value: 'fix', name: 'fix: Bug fix'},
    {value: 'docs', name: 'docs: Documentation changes'},
    {value: 'style', name: 'style: Code style (non-functional changes)'},
    {value: 'refactor', name: 'refactor: Refactoring (no feature/bug fixes)'},
    {value: 'perf', name: 'perf: Performance optimization'},
    {value: 'test', name: 'test: Adding tests'},
    {value: 'chore', name: 'chore: Build or tooling changes'},
    {value: 'revert', name: 'revert: Revert changes'},
    {value: 'build', name: 'build: Packaging'}
  ],
  messages: {
    type: 'Select the type of change:',
    customScope: 'Enter the scope of changes (optional):',
    subject: 'Provide a brief description (required):',
    confirmCommit: 'Confirm commit with the above details? (y/n/e/h)'
  },
  allowCustomScopes: true,
  skipQuestions: ['body', 'footer'],
  subjectLimit: 72
};
```

---

## Automatic Changelog Generation

### Install Conventional Changelog

```bash
npm i conventional-changelog-cli --save-dev
```

Add the following script to `package.json`:

```json
"scripts": {
  "changelog": "conventional-changelog -p angular -i CHANGELOG.md -s"
}
```

Generate the `CHANGELOG.md` file:

```bash
npm run changelog
```

---

## Using the Configuration

1. Merge the tested configuration into the main branch.
2. Other developers can install dependencies with:

   ```bash
   npm install
   ```

3. To use Commitizen, install it globally:

   ```bash
   npm install -g commitizen
   ```

---

## FAQs

### How to merge multiple commits for a single feature in the Changelog?

Currently, automatic merging is not supported. The following steps are recommended:

- Use **Git rebase** to clean up commit history, ensuring one commit per feature.
- Use a PR workflow to enforce rebased commits before merging.

---

### Can I manually select the Commit type?

Yes, Commitizen provides an interactive command `git cz` for selecting commit types. However, Commitlint and Husky are required to enforce validation.

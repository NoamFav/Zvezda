# ⭐ Zvezda

<div align="center">

<img src="https://img.shields.io/badge/go-1.15+-00ADD8.svg?style=for-the-badge&logo=go" alt="Go">
<img src="https://img.shields.io/badge/python-3.8+-3776AB.svg?style=for-the-badge&logo=python" alt="Python">
<img src="https://img.shields.io/badge/license-MIT-green.svg?style=for-the-badge" alt="License">

**Multi-repository management with AI-powered commit generation**

[Installation](#installation) · [Usage](#usage) · [Features](#features)

</div>

---

Zvezda is the original Python + Go repository manager — the predecessor to [Iskra](../Iskra). It provides AI commit generation, batch repo operations, and GitHub clone/sync across all your projects.

---

## Components

| Component | Language | Purpose |
|-----------|----------|---------|
| `ai_commit` | Go | AI-powered commit message generation |
| `auto_commit` | Python | Batch commit across multiple repos |
| `pull_repos` | Python | Bulk clone from GitHub |

---

## Installation

```bash
git clone https://github.com/NoamFav/Zvezda
cd Zvezda

# Install Python dependencies
pip install -r requirements.txt

# Build Go binary
cd src/ai_commit && go build -o ai_commit
sudo mv ai_commit /usr/local/bin/
```

---

## Usage

```bash
# AI commit for current repo
ai_commit

# Batch commit all repos
auto_commit

# Clone all your GitHub repos
pull_repos

# Commit only a specific repo
auto_commit --only my-project

# Pull before committing
auto_commit --pull
```

---

## Features

- AI commit messages via Ollama (Mistral)
- Batch processing across multiple repos
- GitHub repository clone and sync
- Rich terminal UI with progress tracking
- Conventional commit format

---

## License

MIT — see [LICENSE](LICENSE).

---

<div align="center">
Made with ❤️ by <a href="https://github.com/NoamFav">NoamFav</a>
</div>

<div align="center">
<picture>
  <source media="(prefers-color-scheme: dark)" srcset=".github/brand/banner-night.svg">
  <img alt="Zvezda: A comprehensive toolkit for managing multiple Git repositories with intelligent AI-powered commit messages, automatic syncing, and GitHub repository cloning." src=".github/brand/banner-paper.svg" width="100%">
</picture>
<br><br>
<a href="#about"><picture><source media="(prefers-color-scheme: dark)" srcset=".github/brand/tab-about-night.svg"><img alt="about" src=".github/brand/tab-about-paper.svg"></picture></a>
<a href="#components"><picture><source media="(prefers-color-scheme: dark)" srcset=".github/brand/tab-components-night.svg"><img alt="components" src=".github/brand/tab-components-paper.svg"></picture></a>
<a href="#quickstart"><picture><source media="(prefers-color-scheme: dark)" srcset=".github/brand/tab-quickstart-night.svg"><img alt="quickstart" src=".github/brand/tab-quickstart-paper.svg"></picture></a>
</div>

<p>
<a name="about"></a>
<picture><source media="(prefers-color-scheme: dark)" srcset=".github/brand/section-about-night.svg"><img alt="about" src=".github/brand/section-about-paper.svg" width="100%"></picture>
</p>

Zvezda is a multi-repository Git toolkit: it commits, syncs, and clones across every repo you own without you typing `git add . && git commit -m "..."` fifteen times a day. A local LLM (Ollama / Mistral) writes the commit messages; Zvezda handles the batching.

It's the original Python + Go repo manager — the predecessor to [Iskra](https://github.com/NoamFav/Iskra).

<p>
<a name="components"></a>
<picture><source media="(prefers-color-scheme: dark)" srcset=".github/brand/section-components-night.svg"><img alt="components" src=".github/brand/section-components-paper.svg" width="100%"></picture>
</p>

| Component | Language | Purpose |
|-----------|----------|---------|
| `ai_commit` | Go | AI-generated commit messages via Ollama |
| `auto_commit` | Python | Batch commit across multiple repos |
| `pull_repos` | Python | Bulk clone from GitHub |

**Features:** conventional commit format · rich terminal UI with progress tracking · per-repo filtering (`--only`) · pull-before-commit (`--pull`)

<p>
<a name="quickstart"></a>
<picture><source media="(prefers-color-scheme: dark)" srcset=".github/brand/section-quickstart-night.svg"><img alt="quickstart" src=".github/brand/section-quickstart-paper.svg" width="100%"></picture>
</p>

```sh
git clone https://github.com/NoamFav/Zvezda && cd Zvezda

pip install -r requirements.txt
cd src/ai_commit && go build -o ai_commit && sudo mv ai_commit /usr/local/bin/

# AI commit for the current repo
ai_commit

# Batch commit every repo
auto_commit

# Clone all your GitHub repos
pull_repos

# Scoped / with a pull first
auto_commit --only my-project
auto_commit --pull
```

<div align="center">

Made with ♥ by [NoamFav](https://github.com/NoamFav) · MIT License

</div>

<br>

<a href="https://nf-software.com">
<picture>
  <source media="(prefers-color-scheme: dark)" srcset=".github/brand/footer-night.svg">
  <img alt="NF Software" src=".github/brand/footer-paper.svg" width="100%">
</picture>
</a>

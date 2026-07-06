<div align="center">

<img src="https://capsule-render.vercel.app/api?type=venom&height=220&color=gradient&customColorList=12&text=ZVEZDA&fontSize=90&fontColor=fff&animation=twinkling&desc=AI%20Git%20Ops%20for%20Every%20Repo%20You%20Own&descSize=18&descAlignY=65&stroke=FFFFFF&strokeWidth=1" alt="Zvezda Banner" />

<img src="https://readme-typing-svg.herokuapp.com?font=Fira+Code&size=20&pause=1000&color=00D9FF&center=true&vCenter=true&multiline=true&repeat=true&width=900&height=60&lines=Batch+commit+%C2%B7+AI+messages+%C2%B7+clone%2Fsync+every+repo;Go+%2B+Python+%C2%B7+Ollama-powered" alt="Typing SVG" />

<br>

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white&labelColor=0D1117)](https://go.dev)
[![Python](https://img.shields.io/badge/Python-3776AB?style=for-the-badge&logo=python&logoColor=white&labelColor=0D1117)](https://python.org)
[![License](https://img.shields.io/badge/MIT-00D9FF?style=for-the-badge&labelColor=0D1117)](./LICENSE)

</div>

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif" width="100%">

<div align="center">
  <img src="https://readme-typing-svg.herokuapp.com?font=Orbitron&size=26&pause=1000&color=00D9FF&center=true&width=800&lines=%F0%9F%A4%96+WHAT+IS+ZVEZDA+%3F" alt="What is Zvezda" />
</div>
<br>

Zvezda is a multi-repository Git toolkit: it commits, syncs, and clones across every repo you own without you typing `git add . && git commit -m "..."` fifteen times a day. A local LLM (Ollama / Mistral) writes the commit messages; Zvezda handles the batching.

It's the original Python + Go repo manager — the predecessor to [Iskra](https://github.com/NoamFav/Iskra).

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif" width="100%">

<div align="center">
  <img src="https://readme-typing-svg.herokuapp.com?font=Orbitron&size=26&pause=1000&color=FF69B4&center=true&width=800&lines=%F0%9F%9A%80+COMPONENTS+%F0%9F%9A%80" alt="Components" />
</div>
<br>

| Component | Language | Purpose |
|-----------|----------|---------|
| `ai_commit` | Go | AI-generated commit messages via Ollama |
| `auto_commit` | Python | Batch commit across multiple repos |
| `pull_repos` | Python | Bulk clone from GitHub |

**Features:** conventional commit format · rich terminal UI with progress tracking · per-repo filtering (`--only`) · pull-before-commit (`--pull`)

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif" width="100%">

<div align="center">
  <img src="https://readme-typing-svg.herokuapp.com?font=Orbitron&size=26&pause=1000&color=6A5ACD&center=true&width=800&lines=%E2%9C%A8+QUICKSTART+%E2%9C%A8" alt="Quickstart" />
</div>
<br>

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

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif" width="100%">

<div align="center">

<img src="https://readme-typing-svg.herokuapp.com?font=Orbitron&size=20&pause=1000&color=6A5ACD&center=true&width=800&lines=Thanks+for+stopping+by!;Feel+free+to+%E2%AD%90+the+repo." alt="Footer typing" />

<br>

Made with ♥ by [NoamFav](https://github.com/NoamFav) · MIT License

<img src="https://capsule-render.vercel.app/api?type=waving&height=100&color=gradient&customColorList=12&section=footer" />

</div>

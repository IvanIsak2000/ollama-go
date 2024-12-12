# ollama-go

# Basic 
Handler for local `ollama` server for generate an answer from prompt.

# Usage
1. Download ollama:
```
curl -fsSL https://ollama.com/install.sh | sh
```
2. Run ollama server:
```
ollama serve
```
3. Pull `ollama` model, e.g.
```
ollama pull llama3.2:1b
```
4. Exit `ollama serve`
5. Start main.go

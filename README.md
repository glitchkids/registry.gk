# registry.gk

CLI Go pour indexer des dossiers en JSON statique et récupérer des templates depuis un registre distant — sans `git sparse-checkout`.

## Pourquoi

- Démarrer un nouveau projet à partir de templates, en une commande
- Éviter la complexité de `git sparse-checkout`
- Héberger un registre n'importe où en statique (GitHub Pages, CDN, repo Git…) — ce ne sont que des fichiers JSON
- Garder vos habitudes de structure, même avec l'aide de l'IA
- Apprendre Go en construisant un outil utile

## Prérequis

- [Go](https://go.dev/dl/) 1.26.1 ou plus récent

## Installation

```bash
go install glitchkids/registry_gk@latest
```

Ou, depuis une copie locale du dépôt :

```bash
git clone https://github.com/glitchkids/registry.gk
cd registry.gk
go build -o registry_gk .
```

Le binaire s'appelle `registry_gk`.

## Vue d'ensemble

```
┌─────────────────────┐     index      ┌──────────────────┐
│  registry.config    │ ──────────────► │  .registry/      │
│  + vos templates    │                 │  index.json + …  │
└─────────────────────┘                 └────────┬─────────┘
                                                 │ déployer (statique)
                                                 ▼
                                        ┌──────────────────┐
                                        │  URL publique    │
                                        └────────┬─────────┘
                                                 │ add + pull
                                                 ▼
                                        ┌──────────────────┐
                                        │  votre projet    │
                                        └──────────────────┘
```

1. **`index`** — parcourt les dossiers déclarés dans `registry.config.json` et génère un registre JSON local.
2. **Déploiement** — publiez le dossier de sortie (par ex. `.registry/`) sur une URL statique.
3. **`add`** — enregistre l'URL de base du registre distant dans votre config utilisateur.
4. **`pull`** — liste ou télécharge un template depuis ce registre.

## Commandes

### `index`

Génère le registre à partir de `registry.config.json` dans le répertoire courant.

```bash
registry_gk index
```

### `add`

Ajoute un registre distant. L'URL doit être la **base** du registre, **sans** `/index.json`.

```bash
registry_gk add glitchkids https://example.com/path/to/registry
```

### `ls`

Liste les registres distants enregistrés localement.

```bash
registry_gk ls
```

### `remove`

Supprime un registre distant enregistré.

```bash
registry_gk remove glitchkids
```

### `pull`

Sans nom de template : affiche le contenu disponible sur le registre distant.

```bash
registry_gk pull glitchkids
```

Avec un nom de template : télécharge les fichiers dans le répertoire de sortie.

```bash
registry_gk pull glitchkids "Front/Vite-Svelte-UnoCSS" -o ./mon-projet
```

| Option | Alias | Défaut | Description |
|--------|-------|--------|-------------|
| `--out` | `-o` | `.` | Répertoire de destination |
| `--force` | `-f` | `false` | Écrase les fichiers existants |

## Configuration

### `registry.config.json` (indexation)

Fichier lu par `registry_gk index` dans le répertoire courant.

```jsonc
{
  "output": ".registry", // optionnel, défaut : ".registry"
  "input": {
    "folders": [
      {
        "name": "Modules/Node",
        "path": "./templates/node-package",
        "ignores": {
          "files": [".env"],
          "folders": ["node_modules", "dist"]
        }
      },
      {
        "name": "Front/Vite-Svelte-UnoCSS",
        "path": "./templates/vite-svelte-unocss",
        "ignores": {
          "files": [".env"],
          "folders": ["node_modules", "dist"]
        }
      }
    ],
    "files": [] // non implémenté pour l'instant
  }
}
```

Un exemple complet est disponible dans le dépôt [glitchkids/templates](https://github.com/glitchkids/templates).

| Champ | Description |
|-------|-------------|
| `output` | Dossier de sortie du registre généré (chemin local uniquement) |
| `input.folders[].name` | Nom affiché dans l'index (peut contenir `/`) |
| `input.folders[].path` | Chemin local du dossier à indexer |
| `input.folders[].ignores.files` | Fichiers ou motifs glob (`*`) à exclure |
| `input.folders[].ignores.folders` | Sous-dossiers à exclure |

### Sortie générée

Le dossier de sortie contient :

- **`index.json`** — catalogue des templates : `{ "name", "filename" }`
- **Un fichier JSON par dossier** — nom dérivé du `name` (minuscules, `/` → `_`), ex. `front_vite-svelte-unocss.json`, avec le contenu de chaque fichier indexé

### Registres distants (config utilisateur)

Les registres ajoutés via `add` sont stockés dans :

```
{UserConfigDir}/glitchkids/remote_registry.json
```

Sous Windows : `%AppData%\glitchkids\remote_registry.json`  
Sous macOS : `~/Library/Application Support/glitchkids/remote_registry.json`  
Sous Linux : `~/.config/glitchkids/remote_registry.json`

## Exemple de workflow

```bash
# 1. Créer le registre local
registry_gk index

# 2. Déployer le contenu de .registry/ (GitHub Pages, S3, etc.)

# 3. Enregistrer le registre distant
registry_gk add glitchkids https://glitchkids.github.io/templates

# 4. Voir les templates disponibles
registry_gk pull glitchkids

# 5. Récupérer un template
registry_gk pull glitchkids "Modules/Node" -o ./mon-package
```

## Limitations connues

- L'indexation de fichiers isolés (`input.files`) n'est pas encore implémentée
- `pull` ne gère pour l'instant que les dossiers (pas les entrées fichier individuelles)
- Sans `--force`, le comportement face aux fichiers existants est limité (pas de prompt interactif)

## Licence

Ce projet est sous [licence MIT](LICENSE).

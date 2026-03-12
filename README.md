# registry.gk

Quickly index files and folders and pull content from a remote registry.

## Why 
- Quickly start a new project with templates
- git sparse-checkout made my head burn
- It only json files can be host anywhere statically (especially on github)
- Even if you use AI, you maybe have your habits
- Learn Go

## CLI Features

#### Index
Index current directory and sub-directories based on file a JSON config file.

```jsonc
// registry.config.json
// See a complete example in glitchkids/templates repository

{
  "input": {
    "files": [], // files feature is not implemented but working on it as soon i need it
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
    ]
  }
}


```
And then use the command

```bash
registry_gk index
```

#### Registry

Add remote registry

```bash
registry_gk add glitchkids [url]
```

Add list remote registry

```bash
registry_gk ls
```
Delete remote registry

```bash
registry_gk remove glitchkids
```

List content in remote registry

```bash
registry_gk pull glitchkids
```

Pull content from remote registry

```bash
registry_gk pull glitchkids [folder/file_name] -o ./my-dir
```

## How it works

1. The `index` command build a folder with a index.json and json files that index each input declared in the config files.

2. Deploy your folder online statically or on git repository.

3. Add the remote registry base url (without /index.json) in your config with the CLI.

4. I guess you ready to use it.

## LICENSE

This project is licensed under the [MIT License](LICENSE).
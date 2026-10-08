# Clapclap

This is the Navidrome plugin for the clapclap project.

Main repository: https://github.com/TheRaphael0000/clapclap

# Install from release binary

### 1. Download the binary

```bash
# inside your navidrome plugin folder
wget https://github.com/TheRaphael0000/clapclap_navidrome/releases/latest/download/clapclap.ndp
```

### 2. Navidrome configuration

| In config file  | As an env var      | Value for clapclap |
| --------------- | ------------------ | ------------------ |
| Plugins.Enabled | ND_PLUGINS_ENABLED | true               |
| Agents          | ND_AGENTS          | clapclap           |

Don't forget to restart navidrome to apply the configuration

Navidrome doc: 
- https://www.navidrome.org/docs/usage/configuration/options/
- https://www.navidrome.org/docs/usage/features/plugins/#server-configuration

### 3. Configure and enable the plugin

Set the API_URL to your clapclap server and enable the plugin.

Be sure that your navidrome server can reach the clapclap server if they run in different contained environement for example. You can check navidrome logs to help you for this step. For example: `docker compose logs -f`, then try to start the radio.


# Build from sources

This part help you build the `clapclap.ndp` plugin file

## Build with Docker

```bash
git clone https://github.com/TheRaphael0000/clapclap_navidrome.git
cd clapclap_navidrome
docker compose up
```

## Build locally

requirements:

- tinygo
- go
- zip

```bash
git clone https://github.com/TheRaphael0000/clapclap_navidrome.git
cd clapclap_navidrome
make
```

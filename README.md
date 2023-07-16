# OFDL

`ofdl` is a CLI application written in Golang.

It facilitates authentication, stores metadata in `ofdl.sqlite`, sends downloads to [Aria2](https://aria2.github.io/), and organizes media in [Stash](https://stashapp.cc/).

## Setup

### Scraping

0. Initialize the configuration file at `ofdl.yaml`:
   ```bash
   ofdl config init
   ```
0. Navigate to `chrome://version/` in your favorite flavor of Chromium. Try out [Arc](https://arc.net/gift/23df283b)!
0. Copy the "Executable Path". Paste it in `chromium.exec`:
   ```bash
   ofdl config set chromium.exec "/Applications/Arc.app/Contents/MacOS/Arc"
   ```
0. Copy the "Profile Path". Paste it in `chromium.profile`, but remove the last path segment:
   ```bash
   ofdl config set chromium.profile "~/Library/Application Support/Arc/User Data/"
   ```
0. Exit all instances of Chromium, then run the auth helper:
   ```bash
   ofdl auth
   ```
   Log in, then return to your terminal and press enter. This will extract the session details and appropriately update your `ofdl.yaml` config file. You can now re-open Chromium like normal.

### Downloading

OFDL delegates downloads to Aria2. To run Aria2 in Docker, try this:

```yaml
mkdir downloads
docker run -d \
   --name aria2-pro \
   --restart unless-stopped \
   --log-opt max-size=1m \
   -e PUID=$UID \
   -e PGID=$GID \
   -e UMASK_SET=022 \
   -e RPC_SECRET=secret \
   -e RPC_PORT=6800 \
   -p 6800:6800 \
   -e LISTEN_PORT=6888 \
   -p 6888:6888 \
   -p 6888:6888/udp \
   -v $PWD/downloads:/downloads \
   p3terx/aria2-pro
```

The default config will work out of the box with this container, however you can configure your own aria2 server:

```bash
ofdl config set aria2.address ws://aria2:6800/jsonrpc
ofdl config set aria2.secret my-super-notasecret
ofdl config set aria2.root /mnt/data
```

### Organizing

If you run your own Stash server, OFDL can assign a Studio and Performer, as well as other post metadata. Just configure your stash server address and a corresponding Studio ID:

```bash
ofdl config set stash.address http://stash:9999/graphql
ofdl config set stash.studio_id 1
```

> Note: API authentication is not yet supported.

## Usage

You can add `--help` to any command for more information. For example,
`ofdl --help`.

You can check database statistics by running `ofdl stats`.

### Scraping

0. Scrape subscriptions:
   ```bash
   ofdl scrape subscriptions
   ```
0. Scrape media posts:
   ```bash
   ofdl scrape media-posts
   ```
0. Or, now that you're familiar with it, just scrape both:
   ```bash
   ofdl scrape
   ```

### Downloading

0. Download up to 1,000 undownloaded media:
   ```bash
   ofdl download
   ```

### Organizing

0. First, organize Performers:
   ```bash
   ofdl stash subscriptions
   ```
0. Next, organize up to 1,000 unorganized Scenes and Images:
   ```bash
   ofdl stash media
   ```
0. Or, now that you're familiar with it, just scrape both:
   ```bash
   ofdl stash
   ```

# Thanks

Thanks to [DIGITALCRIMINALS](https://github.com/DIGITALCRIMINALS) for maintaining dynamic parameters.

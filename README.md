# OFDL

`ofdl` is a CLI application written in Golang.

It facilitates authentication, stores metadata in `ofdl.sqlite`, sends downloads
to [Aria2](https://aria2.github.io/) (or downloads directly to disk), and
organizes media in [Stash](https://stashapp.cc/).

## Setup

### Scraping

0. Initialize the configuration file at `ofdl.yaml`:
   ```bash
   ofdl config init
   ```
   > Any subsequent `odfl config` commands in this documentation can be omitted
   > in favor of editing the `ofdl.yaml` file directly, if you're comfortable
   > editing a YAML file.
0. Navigate to `chrome://version/` in your favorite flavor of Chromium.
0. Copy the "Executable Path". Paste it in `chromium.exec`:
   ```bash
   ofdl config set chromium.exec "/Applications/Brave Browser.app/Contents/MacOS/Brave Browser"
   ```
0. Copy the "Profile Path". Paste it in `chromium.profile`:
   ```bash
   ofdl config set chromium.profile "$HOME/Library/Application Support/BraveSoftware/Brave-Browser/Default"
   ```
0. Exit all instances of Chromium, then run the auth helper:
   ```bash
   ofdl auth
   ```
   Log in, then the browser window will close. This will extract the session
   details and appropriately update your `ofdl.yaml` config file. You can now
   re-open Chromium like normal.

### Downloading

#### Local

OFDL can save downloads directly to disk. You can customize the path where media
is saved to:

```bash
ofdl config set downloads.downloader local
ofdl config set downloads.local.root $HOME/Downloads
```

#### Aria2

OFDL can delegate downloads to Aria2. To run Aria2 in Docker, try this:

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

The default config will work out of the box with this container, however you
can configure your own aria2 server:

```bash
ofdl config set downloads.downloader aria2
ofdl config set downloads.aria2.address ws://aria2:6800/jsonrpc
ofdl config set downloads.aria2.secret my-super-notasecret
ofdl config set downloads.aria2.root /mnt/data
```

If Aria2 is running on Windows, specify the platform and use a Windows-style
root directory:

```bash
ofdl config set downloads.aria2.platform windows
ofdl config set downloads.aria2.root "D:\Downloads"
```

### Organizing

If you run your own Stash server, OFDL can assign a Studio and Performer, as
well as other post metadata. Just configure your stash server address and a
corresponding Studio ID:

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
0. Enable or Disable individual subscriptions:
   ```bash
   ofdl subscriptions
   ```
0. Scrape media posts:
   ```bash
   ofdl scrape media-posts
   ```
0. Scrape messages:
   ```bash
   ofdl scrape messages
   ```
0. Or, now that you're familiar with it, just scrape both:
   ```bash
   ofdl scrape
   ```

### Downloading

0. Optionally adjust your "batch size", which specifies how many undownloaded
   media are queued for download.
   ```bash
   ofdl config set downloads.batch-size 50
   ```
0. Download up to 1,000 undownloaded post media:
   ```bash
   ofdl download media-posts
   ```
0. Download up to 1,000 undownloaded message media:
   ```bash
   ofdl download messages
   ```
0. Or, now that you're familiar with it, download up to 1,000 of each:
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

Thanks to [DIGITALCRIMINALS](https://github.com/DIGITALCRIMINALS) for
maintaining dynamic parameters.

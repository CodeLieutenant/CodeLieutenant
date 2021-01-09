const { readdir, readFile, writeFile } = require('fs').promises;
const { join } = require('path');

const STYLES_FILE = join(__dirname, 'views_copy', '_partials', 'styles.html');
const SCRIPTS_FILE = join(__dirname, 'views_copy', '_partials', 'scripts.html');

(async () => {
  try {
    const files = await readdir('./dist', 'utf-8');

    const [stylesBuffer, scriptsBuffer] = await Promise.all([
      readFile(STYLES_FILE),
      readFile(SCRIPTS_FILE),
    ]);

    let styles = stylesBuffer.toString();
    let scripts = scriptsBuffer.toString();

    files.forEach((item) => {
      if (!item.match(/[a-zA-Z0-9]+\.[a-f0-9]+\.(css|js)/)) {
        return;
      }

      const [name, _, ext] = item.split('.');
      if (ext === 'css') {
        styles = styles.replace(`${name}.${ext}`, item);
        return;
      }

      if (ext === 'js') {
        scripts = scripts.replace(`${name}.${ext}`, item);
      }
    });

    await Promise.all([writeFile(STYLES_FILE, styles), writeFile(SCRIPTS_FILE, scripts)]);
  } catch (err) {
    console.error(err);
  }
})();

Tired to lost all contents that you type in textarea (or even worse on WYSIWYG)
on some stupid web-page when it accidentally crashed or you just type CTRL+W?

Want all mega-power that you have in your text editor in every web-page?

Then it's for you.

Usage
=====

Install plugin from the webstore that will redirect edits via HTTP.
TextAid, for example, serves well:

    https://chrome.google.com/webstore/detail/textaid/ppoadiihggafnhokfkpphojggcdigllp

Configure plugin to use http://localhost:8888/ as server and select hotkey as
you like.

Next, `go get github.com/seletskiy/chrome-editor`.

Then, launch `chrome-editor`:
```
chrome-editor <editor-of-your-choice>
```

*Note*: to make things work `<editor of your choice>` must accept filename as
argument.

You can just add this command to the `~/.xinitrc`, so it get launched everytime
you start X session.


Useful moments
==============

To make it uber-useful, couple things like i3-wm and vim:
```
chrome-editor sh -c 'i3-msg "workspace 2" && vim --remote-tab-wait $0 && i3-msg "workspace 2"'
```

Of course, you can move that silly stuff to some binary and run chrome-editor:
```
chrome-editor chrome-to-vim
```

In that case when you press hotkey in browser, textarea contents will be opened
in the remote vim and workspace will be switched to editor (I always have vim
opened on workspace "2"). After you finish your edits, you just type :w|bw, and
got focus switched back to textarea in browser.

*Note*: to make it work `workspace_auto_back_and_forth yes` option should be in
the `~/.i3/config` file, and vim should be started as `vim --servername Vim`.

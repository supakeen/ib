# ib

`ib` is a command line client for [Red Hat's Image Builder](https://console.redhat.com/insights/image-builder). A hosted service that lets you build customized operating system images for various architectures and in various output formats.

## API

To use `ib` you need an offline token which requires you to have a Red Hat Subscription. A free developer subscription gives you access.

To obtain an offline token you visit the [Red Hat API Tokens](https://access.redhat.com/management/api) page and store it in your environment under the name `REDHAT_OFFLINE_TOKEN`.

If you forget to do so then `ib` will tell you the same.

## Examples

This example builds a CentOS Stream 9 QCOW image for the x86 architecture and writes it to a file called `image.qcow`.

```sh
€ ib c -a x86_64 -d centos-9 -t guest-image -o image.qcow
queued 968d1756-9121-47cb-8f27-3582639206aa
waiting ........................
downloading 968d1756-9121-47cb-8f27-3582639206aa
```

If you wish for a bit more control then the queuing, status, and downloading steps can be performed separately with the `ib c q`, `ib c s`, and `ib c d` commands. Check their `-help` for any arguments you might need.

If you want to list available distributions you use `ib i d`, if you want to know which architectures support which image types you can use `ib i a -d centos-9`.

---
type: "tv-episode"
title: "{{ .Title }}"
description: "{{ .Title }}"
episode: "{{ .FileName }}"
aliases: ["/{{ .FileName }}"]
publishdate: "{{ .PublishDate }}"
date: "{{ .Date }}"
episode_image: "{{ .Image }}"
images: ["{{ .Image }}"]
youtube: "{{ .VideoID }}"
---

{{ .Description }}


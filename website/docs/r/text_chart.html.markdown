---
layout: "signalfx"
page_title: "SignalFx: signalfx_text_chart"
sidebar_current: "docs-signalfx-resource-text-chart"
description: |-
  Allows Terraform to create and manage SignalFx text notes
---

# Resource: signalfx_text_chart

This special type of chart doesn’t display any metric data. Rather, it lets you place a text note on the dashboard.

## Example Usage

```tf
resource "signalfx_text_chart" "mynote0" {
  name        = "Important Dashboard Note"
  description = "Lorem ipsum dolor sit amet, laudem tibique iracundia at mea. Nam posse dolores ex, nec cu adhuc putent honestatis"

  markdown = <<-EOF
    1. First ordered list item
    2. Another item
      * Unordered sub-list.
    1. Actual numbers don't matter, just that it's a number
      1. Ordered sub-list
    4. And another item.

       You can have properly indented paragraphs within list items. Notice the blank line above, and the leading spaces (at least one, but we'll use three here to also align the raw Markdown).

       To have a line break without a paragraph, you will need to use two trailing spaces.⋅⋅
       Note that this line is separate, but within the same paragraph.⋅⋅
       (This is contrary to the typical GFM line break behaviour, where trailing spaces are not required.)

    * Unordered list can use asterisks
    - Or minuses
    + Or pluses
    EOF
}
```

## Argument Reference

The following arguments are supported in the resource block:

* `name` - (Required) Name of the text note.
* `markdown` - (Required) Markdown text to display.
* `description` - (Optional) Description of the text note.

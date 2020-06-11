---
layout: "signalfx"
page_title: "SignalFx: signalfx_dashboard"
sidebar_current: "docs-signalfx-resource-dashboard"
description: |-
  Allows Terraform to create and manage SignalFx Dashboards
---

# Resource: signalfx_dashboard

A dashboard is a curated collection of specific charts and supports dimensional [filters](http://docs.signalfx.com/en/latest/dashboards/dashboard-filter-dynamic.html#filter-dashboard-charts), [dashboard variables](http://docs.signalfx.com/en/latest/dashboards/dashboard-filter-dynamic.html#dashboard-variables) and [time range](http://docs.signalfx.com/en/latest/_sidebars-and-includes/using-time-range-selector.html#time-range-selector) options. These options are applied to all charts in the dashboard, providing a consistent view of the data displayed in that dashboard. This also means that when you open a chart to drill down for more details, you are viewing the same data that is visible in the dashboard view.

~> **NOTE** Since every dashboard is included in a `dashboard group` (SignalFx collection of dashboards), you need to create that first and reference it as shown in the example.

## Example Usage

```
resource "signalfx_dashboard" "mydashboard0" {
  name            = "My Dashboard"
  dashboard_group = signalfx_dashboard_group.mydashboardgroup0.id

  time_range = "-30m"

  filter {
    property = "collector"
    values   = ["cpu", "Diamond"]
  }
  variable {
    property = "region"
    alias    = "region"
    values   = ["uswest-1-"]
  }
  chart {
    chart_id = signalfx_time_chart.mychart0.id
    width    = 12
    height   = 1
  }
  chart {
    chart_id = signalfx_time_chart.mychart1.id
    width    = 5
    height   = 2
  }
}
```

**Every SignalFx dashboard is shown as a grid of 12 columns and potentially infinite number of rows.** The dimension of the single column depends on the screen resolution.

When you define a dashboard resource, you need to specify which charts (by `chart_id`) should be displayed in the dashboard, along with layout information determining where on the dashboard the charts should be displayed. You have to assign to every chart a **width** in terms of number of columns to cover up (from 1 to 12) and a **height** in terms of number of rows (more or equal than 1). You can also assign a position in the dashboard grid where you like the graph to stay. In order to do that, you assign a **row** that represents the topmost row of the chart and a **column** that represent the leftmost column of the chart. If by mistake, you wrote a configuration where there are not enough columns to accommodate your charts in a specific row, they will be split in different rows. In case a **row** was specified with value higher than 1, if all the rows above are not filled by other charts, the chart will be placed the **first empty row**.

The are a bunch of use cases where this layout makes things too verbose and hard to work with loops. For those you can now use one of these two layouts: grids and columns.

~> **WARNING** These other layouts are not supported by the SignalFx API and are purely provider-side constructs. As such the provider cannot import them and cannot properly reconcile API-side changes. In other words, if someone changes the charts in the UI it will not be reconciled at the next apply. Also, you may only use one of `chart`, `column`, or `grid` when laying out dashboards. You can, however, use multiple instances of each (e.g. multiple `grid`s) for fancy layout.

### Grid

The dashboard is divided into equal-sized charts (defined by `width` and `height`). If a chart does not fit in the same row (because the total width > max allowed by the dashboard), this and the next ones will be place in the next row(s).

```
resource "signalfx_dashboard" "grid_example" {
    name = "Grid"
    dashboard_group = "${signalfx_dashboard_group.example.id}"
    time_range = "-15m"

    grid {
        chart_ids = ["${concat(signalfx_time_chart.rps.*.id,
                signalfx_time_chart.50ths.*.id,
                signalfx_time_chart.99ths.*.id,
                signalfx_time_chart.idle_workers.*.id,
                signalfx_time_chart.cpu_idle.*.id)}"]
        width = 3
        height = 1
    }
}
```

### Column

The dashboard is divided into equal-sized charts (defined by `width` and `height`). The charts are placed in the grid by column (column number is called `column`).

```
resource "signalfx_dashboard" "load" {
    name = "Load"
    dashboard_group = "${signalfx_dashboard_group.example.id}"

    column {
        chart_ids = ["${signalfx_single_value_chart.rps.*.id}"]
        width = 2
    }
    column {
        chart_ids = ["${signalfx_time_chart.cpu_capacity.*.id}"]
        column = 2
        width = 4
    }
}
```

## Argument Reference

The following arguments are supported in the resource block:

* `name` - (Required) Name of the dashboard.
* `dashboard_group` - (Required) The ID of the dashboard group that contains the dashboard.
* `description` - (Optional) Description of the dashboard.
* `charts_resolution` - (Optional) Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
* `time_range` - (Optional) The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
* `start_time` - (Optional) Seconds since epoch. Used for visualization. You must specify time_span_type = `"absolute"` too.
* `end_time` - (Optional) Seconds since epoch. Used for visualization. You must specify time_span_type = `"absolute"` too.
* `filter` - (Optional) Filter to apply to the charts when displaying the dashboard.
    * `property` - (Required) A metric time series dimension or property name.
    * `negated` - (Optional) Whether this filter should be a not filter. `false` by default.
    * `values` - (Required) List of of strings (which will be treated as an OR filter on the property).
    * `apply_if_exist` - (Optional) If true, this filter will also match data that doesn't have this property at all.
* `variable` - (Optional) Dashboard variable to apply to each chart in the dashboard.
    * `property` - (Required) A metric time series dimension or property name.
    * `alias` - (Required) An alias for the dashboard variable. This text will appear as the label for the dropdown field on the dashboard.
    * `description` - (Optional) Variable description.
    * `values` - (Optional) List of of strings (which will be treated as an OR filter on the property).
    * `value_required` - (Optional) Determines whether a value is required for this variable (and therefore whether it will be possible to view this dashboard without this filter applied). `false` by default.
    * `values_suggested` - (Optional) A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.
    * `restricted_suggestions` - (Optional) If `true`, this variable may only be set to the values listed in `values_suggested` and only these values will appear in autosuggestion menus. `false` by default.
    * `replace_only` - (Optional) If `true`, this variable will only apply to charts that have a filter for the property.
    * `apply_if_exist` - (Optional) If true, this variable will also match data that doesn't have this property at all.
* `chart` - (Optional) Chart ID and layout information for the charts in the dashboard.
    * `chart_id` - (Required) ID of the chart to display.
    * `width` - (Optional) How many columns (out of a total of 12) the chart should take up (between `1` and `12`). `12` by default.
    * `height` - (Optional) How many rows the chart should take up (greater than or equal to `1`). `1` by default.
    * `row` - (Optional) The row to show the chart in (zero-based); if `height > 1`, this value represents the topmost row of the chart (greater than or equal to `0`).
    * `column` - (Optional) The column to show the chart in (zero-based); this value always represents the leftmost column of the chart (between `0` and `11`).
* `grid` - (Optional) Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
    * `chart_ids` - (Required) List of IDs of the charts to display.
    * `width` - (Optional) How many columns (out of a total of 12) every chart should take up (between `1` and `12`). `12` by default.
    * `height` - (Optional) How many rows every chart should take up (greater than or equal to `1`). `1` by default.
* `column` - (Optional) Column layout. Charts listed will be placed in a single column with the same width and height.
    * `chart_ids` - (Required) List of IDs of the charts to display.
    * `column` - (Optional) Column number for the layout.
    * `width` - (Optional) How many columns (out of a total of `12`) every chart should take up (between `1` and `12`). `12` by default.
    * `height` - (Optional) How many rows every chart should take up (greater than or equal to 1). 1 by default.
* `event_overlay` - (Optional) Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selected_event_overlay` property instead.
    * `line` - (Optional) Show a vertical line for the event. `false` by default.
    * `label` - (Optional) Text shown in the dropdown when selecting this overlay from the menu.
    * `color` - (Optional) Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
    * `signal` - Search term used to choose the events shown in the overlay.
    * `type` - (Optional) Can be set to `eventTimeSeries` (the default) to refer to externally reported events, or `detectorEvents` to refer to events from detector triggers.
    * `source` - (Optional) Each element specifies a filter to use against the signal specified in the `signal`.
      * `property` - The name of a dimension to filter against.
      * `values` - A list of values to be used with the `property`, they will be combined via `OR`.
      * `negated` - (Optional) If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
* `selected_event_overlay` - (Optional) Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `event_overlay`, which are similar to the properties here.
    * `signal` - Search term used to choose the events shown in the overlay.
    * `type` - (Optional) Can be set to `eventTimeSeries` (the default) to refer to externally reported events, or `detectorEvents` to refer to events from detector triggers.
    * `source` - (Optional) Each element specifies a filter to use against the signal specified in the `signal`.
      * `property` - The name of a dimension to filter against.
      * `values` - A list of values to be used with the `property`, they will be combined via `OR`.
      * `negated` - (Optional) If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.


## Dashboard Layout Information

**Every SignalFx dashboard is shown as a grid of 12 columns and potentially infinite number of rows.** The dimension of the single column depends on the screen resolution.

When you define a dashboard resource, you need to specify which charts (by `chart_id`) should be displayed in the dashboard, along with layout information determining where on the dashboard the charts should be displayed. You have to assign to every chart a **width** in terms of number of columns to cover up (from 1 to 12) and a **height** in terms of number of rows (more or equal than 1). You can also assign a position in the dashboard grid where you like the graph to stay. In order to do that, you assign a **row** that represents the topmost row of the chart and a **column** that represent the leftmost column of the chart. If by mistake, you wrote a configuration where there are not enough columns to accommodate your charts in a specific row, they will be split in different rows. In case a **row** was specified with value higher than 1, if all the rows above are not filled by other charts, the chart will be placed the **first empty row**.

The are a bunch of use cases where this layout makes things too verbose and hard to work with loops. For those you can now use one of these two layouts: grids and columns.

~> **WARNING** These other layouts are not supported by the SignalFx API and are purely Terraform-side constructs. As such the provider cannot import them and cannot properly reconcile API-side changes. In other words, if someone changes the charts in the UI it will not be reconciled at the next apply. Also, you may only use one of `chart`, `column`, or `grid` when laying out dashboards. You can, however, use multiple instances of each (e.g. multiple `grid`s) for fancy layout.

### Grid

The dashboard is divided into equal-sized charts (defined by `width` and `height`). If a chart does not fit in the same row (because the total width > max allowed by the dashboard), this and the next ones will be place in the next row(s).

```tf
resource "signalfx_dashboard" "grid_example" {
  name            = "Grid"
  dashboard_group = signalfx_dashboard_group.example.id
  time_range      = "-15m"

  grid {
    chart_ids = [
      concat(
        signalfx_time_chart.rps.*.id,
        signalfx_time_chart.p50ths.*.id,
        signalfx_time_chart.p99ths.*.id,
        signalfx_time_chart.idle_workers.*.id,
        signalfx_time_chart.cpu_idle.*.id,
      )
    ]
    width  = 3
    height = 1
  }
}
```

### Column

The dashboard is divided into equal-sized charts (defined by `width` and `height`). The charts are placed in the grid by column (column number is called `column`).

```tf
resource "signalfx_dashboard" "load" {
  name            = "Load"
  dashboard_group = signalfx_dashboard_group.example.id

  column {
    chart_ids = [signalfx_single_value_chart.rps.*.id]
    width     = 2
  }
  column {
    chart_ids = [signalfx_time_chart.cpu_capacity.*.id]
    column    = 2
    width     = 4
  }
}
```

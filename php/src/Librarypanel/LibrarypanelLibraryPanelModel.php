<?php

namespace Grafana\Foundation\Librarypanel;

class LibrarypanelLibraryPanelModel implements \JsonSerializable
{
    /**
     * The panel plugin type id. This is used to find the plugin to display the panel.
     */
    public string $type;

    /**
     * The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
     */
    public ?string $pluginVersion;

    /**
     * Tags for the panel.
     * @var array<string>|null
     */
    public ?array $tags;

    /**
     * Depends on the panel plugin. See the plugin documentation for details.
     * @var array<\Grafana\Foundation\Cog\Dataquery>|null
     */
    public ?array $targets;

    /**
     * Panel title.
     */
    public ?string $title;

    /**
     * Panel description.
     */
    public ?string $description;

    /**
     * Whether to display the panel without a background.
     */
    public bool $transparent;

    /**
     * The datasource used in all targets.
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * Panel links.
     * @var array<\Grafana\Foundation\Dashboard\DashboardLink>|null
     */
    public ?array $links;

    /**
     * Name of template variable to repeat for.
     */
    public ?string $repeat;

    /**
     * Direction to repeat in if 'repeat' is set.
     * `h` for horizontal, `v` for vertical.
     */
    public ?\Grafana\Foundation\Librarypanel\LibraryPanelRepeatDirection $repeatDirection;

    /**
     * Id of the repeating panel.
     */
    public ?int $repeatPanelId;

    /**
     * The maximum number of data points that the panel queries are retrieving.
     */
    public ?float $maxDataPoints;

    /**
     * List of transformations that are applied to the panel data before rendering.
     * When there are multiple transformations, Grafana applies them in the order they are listed.
     * Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
     * @var array<\Grafana\Foundation\Dashboard\DataTransformerConfig>
     */
    public array $transformations;

    /**
     * The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
     * This value must be formatted as a number followed by a valid time
     * identifier like: "40s", "3d", etc.
     * See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
     */
    public ?string $interval;

    /**
     * Overrides the relative time range for individual panels,
     * which causes them to be different than what is selected in
     * the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
     * time periods or days on the same dashboard.
     * The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
     * `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
     * Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
     * See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
     */
    public ?string $timeFrom;

    /**
     * Overrides the time range for individual panels by shifting its start and end relative to the time picker.
     * For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
     * Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
     * See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
     */
    public ?string $timeShift;

    /**
     * It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
     * @var mixed
     */
    public $options;

    /**
     * Field options allow you to change how the data is displayed in your visualizations.
     */
    public \Grafana\Foundation\Dashboard\FieldConfigSource $fieldConfig;

    /**
     * @param string|null $type
     * @param string|null $pluginVersion
     * @param array<string>|null $tags
     * @param array<\Grafana\Foundation\Cog\Dataquery>|null $targets
     * @param string|null $title
     * @param string|null $description
     * @param bool|null $transparent
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     * @param array<\Grafana\Foundation\Dashboard\DashboardLink>|null $links
     * @param string|null $repeat
     * @param \Grafana\Foundation\Librarypanel\LibraryPanelRepeatDirection|null $repeatDirection
     * @param int|null $repeatPanelId
     * @param float|null $maxDataPoints
     * @param array<\Grafana\Foundation\Dashboard\DataTransformerConfig>|null $transformations
     * @param string|null $interval
     * @param string|null $timeFrom
     * @param string|null $timeShift
     * @param mixed|null $options
     * @param \Grafana\Foundation\Dashboard\FieldConfigSource|null $fieldConfig
     */
    public function __construct(?string $type = null, ?string $pluginVersion = null, ?array $tags = null, ?array $targets = null, ?string $title = null, ?string $description = null, ?bool $transparent = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null, ?array $links = null, ?string $repeat = null, ?\Grafana\Foundation\Librarypanel\LibraryPanelRepeatDirection $repeatDirection = null, ?int $repeatPanelId = null, ?float $maxDataPoints = null, ?array $transformations = null, ?string $interval = null, ?string $timeFrom = null, ?string $timeShift = null,  $options = null, ?\Grafana\Foundation\Dashboard\FieldConfigSource $fieldConfig = null)
    {
        $this->type = $type ?: "";
        $this->pluginVersion = $pluginVersion;
        $this->tags = $tags;
        $this->targets = $targets;
        $this->title = $title;
        $this->description = $description;
        $this->transparent = $transparent ?: false;
        $this->datasource = $datasource;
        $this->links = $links;
        $this->repeat = $repeat;
        $this->repeatDirection = $repeatDirection;
        $this->repeatPanelId = $repeatPanelId;
        $this->maxDataPoints = $maxDataPoints;
        $this->transformations = $transformations ?: [];
        $this->interval = $interval;
        $this->timeFrom = $timeFrom;
        $this->timeShift = $timeShift;
        $this->options = $options ?: null;
        $this->fieldConfig = $fieldConfig ?: new \Grafana\Foundation\Dashboard\FieldConfigSource();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, pluginVersion?: string, tags?: array<string>, targets?: array<mixed>, title?: string, description?: string, transparent?: bool, datasource?: mixed, links?: array<mixed>, repeat?: string, repeatDirection?: string, repeatPanelId?: int, maxDataPoints?: float, transformations?: array<mixed>, interval?: string, timeFrom?: string, timeShift?: string, options?: mixed, fieldConfig?: mixed} $inputData */
        $data = $inputData;
        return new self(
            type: $data["type"] ?? null,
            pluginVersion: $data["pluginVersion"] ?? null,
            tags: $data["tags"] ?? null,
            targets: isset($data["targets"]) ? (function ($in) {
    	/** @var array{datasource?: array{type?: mixed}} $in */
        $hint = (isset($in["datasource"], $in["datasource"]["type"]) && is_string($in["datasource"]["type"])) ? $in["datasource"]["type"] : "";
        /** @var array<array<string, mixed>> $in */
        return \Grafana\Foundation\Cog\Runtime::get()->dataqueriesFromArray($in, $hint);
    })($data["targets"]): null,
            title: $data["title"] ?? null,
            description: $data["description"] ?? null,
            transparent: $data["transparent"] ?? null,
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DataSourceRef::fromArray($val);
    })($data["datasource"]) : null,
            links: array_filter(array_map((function($input) {
    	/** @var array{title?: string, type?: string, icon?: string, tooltip?: string, url?: string, tags?: array<string>, asDropdown?: bool, targetBlank?: bool, includeVars?: bool, keepTime?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DashboardLink::fromArray($val);
    }), $data["links"] ?? [])),
            repeat: $data["repeat"] ?? null,
            repeatDirection: isset($data["repeatDirection"]) ? (function($input) { return \Grafana\Foundation\Librarypanel\LibraryPanelRepeatDirection::fromValue($input); })($data["repeatDirection"]) : null,
            repeatPanelId: $data["repeatPanelId"] ?? null,
            maxDataPoints: $data["maxDataPoints"] ?? null,
            transformations: array_filter(array_map((function($input) {
    	/** @var array{id?: string, disabled?: bool, filter?: mixed, options?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DataTransformerConfig::fromArray($val);
    }), $data["transformations"] ?? [])),
            interval: $data["interval"] ?? null,
            timeFrom: $data["timeFrom"] ?? null,
            timeShift: $data["timeShift"] ?? null,
            options: $data["options"] ?? null,
            fieldConfig: isset($data["fieldConfig"]) ? (function($input) {
    	/** @var array{defaults?: mixed, overrides?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\FieldConfigSource::fromArray($val);
    })($data["fieldConfig"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
            "transparent" => $this->transparent,
            "transformations" => $this->transformations,
            "options" => $this->options,
            "fieldConfig" => $this->fieldConfig,
        ];
        if (isset($this->pluginVersion)) {
            $data["pluginVersion"] = $this->pluginVersion;
        }
        if (isset($this->tags)) {
            $data["tags"] = $this->tags;
        }
        if (isset($this->targets)) {
            $data["targets"] = $this->targets;
        }
        if (isset($this->title)) {
            $data["title"] = $this->title;
        }
        if (isset($this->description)) {
            $data["description"] = $this->description;
        }
        if (isset($this->datasource)) {
            $data["datasource"] = $this->datasource;
        }
        if (isset($this->links)) {
            $data["links"] = $this->links;
        }
        if (isset($this->repeat)) {
            $data["repeat"] = $this->repeat;
        }
        if (isset($this->repeatDirection)) {
            $data["repeatDirection"] = $this->repeatDirection;
        }
        if (isset($this->repeatPanelId)) {
            $data["repeatPanelId"] = $this->repeatPanelId;
        }
        if (isset($this->maxDataPoints)) {
            $data["maxDataPoints"] = $this->maxDataPoints;
        }
        if (isset($this->interval)) {
            $data["interval"] = $this->interval;
        }
        if (isset($this->timeFrom)) {
            $data["timeFrom"] = $this->timeFrom;
        }
        if (isset($this->timeShift)) {
            $data["timeShift"] = $this->timeShift;
        }
        return $data;
    }
}
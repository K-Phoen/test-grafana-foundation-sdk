<?php

namespace Grafana\Foundation\Loki;

final class SupportingQueryType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, SupportingQueryType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function logsVolume(): self
    {
        if (!isset(self::$instances["logsVolume"])) {
            self::$instances["logsVolume"] = new self("logsVolume");
        }

        return self::$instances["logsVolume"];
    }

    public static function logsSample(): self
    {
        if (!isset(self::$instances["logsSample"])) {
            self::$instances["logsSample"] = new self("logsSample");
        }

        return self::$instances["logsSample"];
    }

    public static function dataSample(): self
    {
        if (!isset(self::$instances["dataSample"])) {
            self::$instances["dataSample"] = new self("dataSample");
        }

        return self::$instances["dataSample"];
    }

    public static function infiniteScroll(): self
    {
        if (!isset(self::$instances["infiniteScroll"])) {
            self::$instances["infiniteScroll"] = new self("infiniteScroll");
        }

        return self::$instances["infiniteScroll"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "logsVolume") {
            return self::logsVolume();
        }

        if ($value === "logsSample") {
            return self::logsSample();
        }

        if ($value === "dataSample") {
            return self::dataSample();
        }

        if ($value === "infiniteScroll") {
            return self::infiniteScroll();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum SupportingQueryType");
    }

    public function jsonSerialize(): string
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return $this->value;
    }
}

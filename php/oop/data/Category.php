<?php

class Category
{
    private string $name;
    private bool $expensive;

    /**
     * @return string
     */
    public function getName(): string
    {
        return $this->name;
    }

    /**
     * @param string $name
     */
    public function setName(string $name): void
    {
        if (trim($name) != ""){
            $this->name = $name;
        }
    }

    /**
     * @return bool
     */
    public function isExpensive(): bool
    {
        return $this->expensive;
    }

    /**
     * @param bool $expensive
     */
    public function setExpensive(bool $expensive): void
    {
        $this->expensive = $expensive;
    }
}
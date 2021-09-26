<?php

namespace Jstmade\Mvc\Middleware;

interface Middleware
{
    function before(): void;
}
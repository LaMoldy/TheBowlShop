<?php

use Illuminate\Support\Facades\Route;

Route::get('/testing', function () {
    return response()->json([
        "message" => "hello, world"
    ]);
});

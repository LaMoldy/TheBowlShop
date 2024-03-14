<?php

namespace App\Http\Controllers;

class UserController extends Controller
{
    public function login(string $email, string $password)
    {
        return view('testing');
    }
}

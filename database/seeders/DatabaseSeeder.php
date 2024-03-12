<?php

namespace Database\Seeders;

use App\Models\User;
use Carbon\Carbon;
use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    public function run(): void
    {
        User::factory(1)->create();

        User::factory()->create([
            'email' => 'test@example.com',
            'first_name' => 'Test',
            'last_name' => 'User',
            'created_at' => Carbon::now(), // Carbon gets the current time
            'updated_at' => Carbon::now(),
            'is_admin' => true,
            'is_deleted' => false
        ]);
    }
}

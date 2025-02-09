<?php

namespace Database\Seeders;

use App\Models\User;
// use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use App\Models\Master\MstRestaurant;
use App\Models\Master\MstCategory;
use Illuminate\Support\Facades\Hash;
use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    /**
     * Seed the application's database.
     */
    public function run(): void
    {
        // User::factory(10)->create();

        User::create([
            'username' => 'Test',
            'password' => Hash::make('password'),
            'email' => 'test@example.com',
        ]);

        MstRestaurant::create([
            'name' => 'Restaurant 1',
            'description' => 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus.',
            'phone_number' => '08123456789',
            'avatar' => '/Images/default-avatar-restaurant.jpg',
            'category_id' => 1
        ]);

        MstRestaurant::create([
            'name' => 'Restaurant 2',
            'description' => 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus.',
            'avatar' => '/Images/default-avatar-restaurant.jpg',
            'phone_number' => '08123456789',
            'category_id' => 2
        ]);

        MstCategory::create([
            'name' => 'Kiri'
        ]);

        MstCategory::create([
            'name' => 'Kanan'
        ]);

    }
}

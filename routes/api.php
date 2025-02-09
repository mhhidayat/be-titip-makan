<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\Auth\AuthController;
use App\Http\Controllers\Restaurant\RestaurantController;

Route::group([
    'prefix' => 'v1',
    'middleware' => 'auth',
], function () {

    Route::group(['prefix' => 'auth'], function () {
        Route::post('/login', [AuthController::class, 'login'])->withoutMiddleware('auth');
        Route::get('/me', [AuthController::class, 'me']);
        Route::post('/logout', [AuthController::class, 'logout']);
        Route::post('/refresh', [AuthController::class, 'refresh']);
    });

    Route::get('/get-restaurants', [RestaurantController::class, 'getRestaurants']);
    Route::get('/get-categories', [RestaurantController::class, 'getCategories']);

});

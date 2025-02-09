<?php

namespace App\Http\Controllers\Restaurant;

use Illuminate\Http\Request;
use App\Models\Master\MstCategory;
use Illuminate\Routing\Controller;
use App\Models\Master\MstRestaurant;

class RestaurantController extends Controller
{

    public function getRestaurants(Request $request)
    {

        $data = MstRestaurant::with('category');

        if ($request->has('category') && is_int( (int) $request->category)) {
            $data->where('category_id', $request->category);
        }

        return res('Data restaurant', $data->get(), 200);
    }

    public function getCategories(Request $request)
    {
        $data = MstCategory::all();
        return res('Data category', $data, 200);
    }

}

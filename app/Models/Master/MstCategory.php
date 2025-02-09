<?php

namespace App\Models\Master;

use Illuminate\Database\Eloquent\Model;

class MstCategory extends Model
{
    public $table = 'mst_categorys';
    public $hidden = [
        'created_at',
        'updated_at'
    ];


    public function restaurant()
    {
        return $this->hasMany(MstRestaurant::class, 'category_id', 'id');
    }

}

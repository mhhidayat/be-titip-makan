<?php

namespace App\Models\Master;

use Illuminate\Database\Eloquent\Model;

class MstRestaurant extends Model
{

    public $table = 'mst_restaurants';
    public $hidden = [
        'created_at',
        'updated_at'
    ];

    protected $appends = ['avatar_url'];

    public function getAvatarUrlAttribute()
    {
        return $this->avatar ? asset('storage/' . $this->avatar) : null;
    }

    public function category()
    {
        return $this->belongsTo(MstCategory::class, 'category_id', 'id');
    }

}

<?php

namespace App\Http\Controllers\Users;

use App\Http\Requests\AuthRequest;
use Illuminate\Routing\Controller;

class AuthController extends Controller
{

    public function __construct()
    {
        $this->middleware('auth', ['except' => ['login']]);
    }

    public function login(AuthRequest $request)
    {
        $credentials = request(['username', 'password']);

        if (! $token = auth()->attempt($credentials)) {
            return res('Username atau password salah', $this->respondWithToken($token), 401);
        }

        return res('Berhasil login', $this->respondWithToken($token), 200);
    }

    public function me()
    {
        return res('Data user', auth()->user(), 200);
    }

    public function logout()
    {
        auth()->logout(true);
        return res('Berhasil logout', '', 200);
    }

    public function refresh()
    {
        return res('Berhasil refresh', $this->respondWithToken(auth()->refresh(true, true)), 200);
    }

    protected function respondWithToken($token)
    {
        return response()->json([
            'access_token' => $token,
            'token_type' => 'bearer',
            'expires_in' => auth()->factory()->getTTL() * 60
        ]);
    }

}

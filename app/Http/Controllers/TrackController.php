<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

use App\Track;
use App\Http\Resources\Track as TrackResource;
use App\Http\Resources\Tracks as TracksResource;


class TrackController extends Controller
{

    /**
     * Show all resources in the database
     *
     * @param  \Illuminate\Http\Request  $request
     * @return App\Http\Resources\Tracks
     */
    public function index()
    {
        return new TracksResource(Track::all());
    }

    /**
     * Show the specified resource
     *
     * @param  App\Track $track
     * @return App\Http\Resources\Track
     */
    public function show(Track $track)
    {
        return new TrackResource($track);
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        $track = Track::create($request->all());
        return TrackResource($track);
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  Track $track
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, Track $track)
    {
        $track->update($request->all());

        return TrackResource($track);
    }


    /**
     * Remove the specified resource from storage.
     *
     * @param  Track $track
     * @return \Illuminate\Http\Response
     */
    public function destroy(Track $track)
    {
        $track->delete();

        return response()->json(null, 204);
    }
}

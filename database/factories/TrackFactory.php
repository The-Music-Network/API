<?php

use Faker\Generator as Faker;

/*
|--------------------------------------------------------------------------
| Model Factories
|--------------------------------------------------------------------------
|
| This directory should contain each of the model factory definitions for
| your application. Factories provide a convenient way to generate new
| model instances for testing / seeding your application's database.
|
*/

$factory->define(App\Track::class, function (Faker $faker) {
    return [
        'title' => $faker->unique()->text(),
        'user_id' => $faker->randomElement([1,2,3,4,5,6,7,8,9,10]),
        'track_url' => $faker->randomElement(['/test/file.mp3', '/test1/file1.mp3', '/test2/file2.mp3', '/test3/file3.mp3', '/test4/file4.mp3']),
    ];
});

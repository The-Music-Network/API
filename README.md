# API
This is the backend API for The-Music-Network

## Resources

<h4>Actions Handled By Resource Controller</h4>
<b>Notice, we define basic CRUD api endpoints as a single resource route in routes/api.php, like so:</b>
<code>Route::resource('/tracks', 'TrackController');</code>

By using the Route::resource, the corresponding controller automatically has the following actions and routes:
<table>
<thead>
<tr>
<th>Verb</th>
<th>URI</th>
<th>Action</th>
<th>Route Name</th>
</tr>
</thead>
<tbody>
<tr>
<td>GET</td>
<td><code>/users</code></td>
<td>index</td>
<td>users.index</td>
</tr>
<tr>
<td>POST</td>
<td><code>/users</code></td>
<td>store</td>
<td>users.store</td>
</tr>
<tr>
<td>PUT/PATCH</td>
<td><code>/users/{photo}</code></td>
<td>update</td>
<td>users.update</td>
</tr>
<tr>
<td>DELETE</td>
<td><code>/users/{photo}</code></td>
<td>destroy</td>
<td>users.destroy</td>
</tr>
</tbody>
</table>

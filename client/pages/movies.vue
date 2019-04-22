<template>
  <v-layout column>
    <v-layout row>
      <v-text-field label="Movies Path" single-line v-model="moviesPath"></v-text-field>
      <v-btn color="primary" @click="sync()">Sync</v-btn>
    </v-layout>
    <v-container fluid grid-list-sm>
      <v-layout row wrap>
        <v-flex v-for="movie in movies" :key="movie" xs2>
          <img :src="movie.image_url" class="image" :alt="movie.title" width="100%" height="100%">
        </v-flex>
      </v-layout>
    </v-container>
  </v-layout>
</template>

<script>
  import * as p from '../pb/requests_pb'
  import { MoviesRequestsClient } from '../pb/requests_grpc_web_pb.js'

  export default {
    data() {
      return {
        moviesPath: "/Users/christopher.ganga/Dev/osc/vida/data",// should be changed
        movies: [],
        proto: null
      }
    },
    methods: {
      sync() {
        // console.log(p);
        this.proto = p;
        this.getMovies()
      },
      getMovies() {
        var vidaClient = new MoviesRequestsClient("http://localhost:50005");

        var r = new proto.Empty();
        var call = vidaClient.listMovies(r, {}, function(err, response) {
          if (err) {
            console.log(err.code);
            console.log(err.message);
          } else {
            console.log(response.getMessage());
          }
        });
        call.on('status', function(status) {
          console.log(status.code);
          console.log(status.details);
          console.log(status.metadata);
        });
      }
    }
  }
</script>

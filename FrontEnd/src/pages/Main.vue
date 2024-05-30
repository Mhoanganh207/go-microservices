<script setup lang="ts">

import axios from 'axios'

async function testBroker(){
   let output = document.getElementById("output")
   let sent = document.getElementById("payload")
   let received = document.getElementById("received")

   const data = await axios.post("http://localhost:8000/broker").then(res => {
      return res.data;
   })
   .catch(err => {
    if(output) {
        output.innerHTML = `<br><strong>Error <strong> : ${err.message}`
      }
   })

  if (sent) {
    sent.innerHTML = 'empty post request';
  }

  if(received){
    received.innerHTML = JSON.stringify(data, undefined,4)
  }

  if(output) {
        output.innerHTML = `<br><strong>Response from broker server<strong> : ${data.message}`
      }


   





}

</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5" style="font-weight: bold;">Test microservices</h1>
        <hr>
        <button @click="testBroker" id="broker-btn" class="btn btn-outline-secondary">Test Broker</button>

        <div id="output" class="mt-5" style="outline: 1px solid silver; padding: 2em; border-radius: 15px;">
          <span class="text-muted">Output shows here...</span>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <h4 class="mt-5" style="font-weight: bold;">Sent</h4>
        <div class="mt-1" style="outline: 1px solid silver; padding: 2em; border-radius: 15px;">
          <pre id="payload"><span class="text-muted">Nothing sent yet...</span></pre>
        </div>
      </div>
      <div class="col">
        <h4 class="mt-5" style="font-weight: bold;">Received</h4>
        <div class="mt-1" style="outline: 1px solid silver; padding: 2em; border-radius: 15px;">
          <pre id="received"><span class="text-muted">Nothing received yet...</span></pre>
        </div>
      </div>
    </div>
  </div>

</template>

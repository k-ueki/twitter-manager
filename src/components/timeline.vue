<template>
	<div>
		<h1>Timeline</h1>
		<button @click="getTimeline">getTimeline</button>
		<div v-for="tl in timeline">{{tl}}<br/><br/></div>
	</div>
</template>
<script>
import axios from 'axios'
	
const base = "http://localhost:7777/timeline/"

export default{
	name:"main",
	data(){
		return {
			timeline:[],
		}
	},
	created(){
		axios.get(base)
		.then(res => {
			this.timeline = res.data
			console.log(res.data)
		}).catch(err => {
			console.log(err)
		})
	},
	methods:{
		getTimeline(){
			axios.get(base)
				.then(response => {
					this.timeline = response.data
					console.log(response.data)
				}).catch(error => {
					console.log(error)
				})
		}
	}
}

</script>
<style>

</style>

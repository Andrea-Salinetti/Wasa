<script>

export default {
	data: function() {
		return {
			username: sessionStorage.getItem('username'),
			userId: sessionStorage.getItem('userId'),
			logged: sessionStorage.getItem('logged'),
			deleteMode: false,
			profile:  [],
			toVisitProfile: [],
			errormsg: null,
			positiveBanner: null,
			profileRetrieved: null

		}
	},
	mounted() {
		this.getUserProfile(sessionStorage.getItem('username'))
	},
	methods: {
		async getUserProfile(username){
			if (!sessionStorage.getItem('logged')){
				return
			}
			this.loading = true;
			this.errormsg = null;
			this.profileRetrieved = false;
			let toVisitUsername;
			if (username=='others') {
				toVisitUsername = document.getElementById('others-profile-input').value} else {
					toVisitUsername = username
				}
			console.log(username)
			const config = {
				headers: {
				'Authorization': 'Bearer ' + sessionStorage.getItem('userId')
				}
			};
			try {
				let response = await this.$axios.get("/profiles/"+ toVisitUsername + "?userId=" + sessionStorage.getItem('userId'), config);
				if (toVisitUsername == username){
					this.profile = response.data
				} else {
					this.toVisitProfile = response.data
					this.profileRetrieved = toVisitUsername
				}

			} catch (e) {
				this.errormsg = e.toString();
			}
			if (this.profile.length){
				const deleteButton = document.getElementById("delete-button")
				deleteButton.disabled = false
			}
			this.loading = false;
		},

		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/");
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			console.log("refreshing")
		},

		async uploadPhoto() {
		console.log('userId ' + sessionStorage.getItem('userId'))
  		const fileInput = document.getElementById('photoInput');
 		const file = fileInput.files[0];
		if (!file) {
			return;
		}

		const formData = new FormData();
		formData.append("image", file)
		try {
			const url = "/photos/?userId=" + sessionStorage.getItem('userId');
			const config = {
				headers: {
				'Authorization': 'Bearer ' + sessionStorage.getItem('userId')
				}
			};
			const response = await this.$axios.post(url, formData, config);
			if (response.status == 201){
				const newPhoto = {
				photoId: response.data
			}
			this.positiveBanner = 'Photo succesfully uploaded'
				setTimeout(()=>{
					this.positiveBanner = null
				}, 3000)
			this.profile.push(newPhoto)
			if (this.profile.length){
				const deleteButton = document.getElementById("delete-button")
				deleteButton.disabled = false
			}
			}

    } catch (error) {
      console.error("Error occurred:", error);
    }
},

	async toggleDeleteMode() {
	if (this.deleteMode == false){
		var deleteButton = document.getElementById('delete-button');
		var photos = document.getElementsByClassName("profile-image")
		var buttons = document.querySelectorAll('button');
		var selector = document.getElementById('photoInput');
		selector.disabled = true

		buttons.forEach(function(button) {
    		button.disabled = true;
		});
		deleteButton.disabled = false
		this.deleteMode = true;
		deleteButton.textContent = 'Cancel';


		for (var i = 0; i < photos.length; i++) {
			photos[i].style.borderColor = "#EE5535";
			photos[i].style.borderWidth = '3px';
			photos[i].style.borderStyle = 'solid';
			// photos[i].style.backgroundColor = "#EE5535";
			photos[i].disabled = false;
	}}
	else{
		var deleteButton = document.getElementById('delete-button');
		this.deleteMode = false;
		deleteButton.textContent = 'Delete Photo';
		var photos = document.getElementsByClassName("profile-image")
		for (var i = 0; i < photos.length; i++) {
			// photos[i].style.backgroundColor = "transparent";
			photos[i].style.border = "none";
			photos[i].disabled = true;
			}
		var buttons = document.querySelectorAll('button');
		var selector = document.getElementById('photoInput');
		selector.disabled = false

		buttons.forEach(function(button) {
    	button.disabled = false;
		if (photos.length == 0){
			deleteButton.disabled = true
		}
		});
}      
},	
async deletePhoto(photoId){
		this.loading = true
		if (this.deleteMode){
		var photo = document.getElementById(photoId)
		this.toggleDeleteMode()
		try{
			const response = this.$axios.delete("/photos/" + photoId + "?userId=" + sessionStorage.getItem('userId'), {
				headers:{
					'Authorization': 'Bearer ' + sessionStorage.getItem('userId')
				}
			})
			
			for (var i = 0; i < this.profile.length; i++) {
				if (this.profile[i].photoId == photoId){
					this.profile.splice(i,1);
					break
				}
			}
			console.log(this.profile)
			photo.remove();
			window.location.reload()
			const deleteButton = document.getElementById("delete-button")
		if (this.profile.length){
			deleteButton.disabled = false
		} else{
			deleteButton.disabled = true
		}
		this.loading = false
		} catch (e) {
				this.errormsg = e.toString();
				console.log(this.errormsg)
			}
		}
		
	}
}}

</script>

<template>
	<p v-if="!this.logged" class="positive-banner" style="margin-top: 40px;">Here you will be able to see your profile, but please log-in first in the home page</p>
	<div v-if="this.logged">
		<div>
			<div
				class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
				<h1 class="h2">Hi {{ this.username }}, this is your profile</h1>
			</div>
			<p v-if="positiveBanner" class="positive-banner">{{ this.positiveBanner }}</p>
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

			<div>
				<div style="display: flex; flex-direction: row; margin-bottom: 20px;">
					<input type="file" class="btn btn-outline-secondary" id="photoInput" ref="photoInput" @change="uploadPhoto" accept="image/*">
					<button class="btn btn-outline-secondary" id="post-button" type="button" style="margin-left: 10px" @click="uploadPhoto">Post Photo</button>
				</div>
				<p v-if="deleteMode" class="h3" style="margin-top: 20px;">Click on the image you want to delete</p>
			</div>
			<div>
				<div id="image-grid" class="image-grid" v-if="this.profile.length && !this.loading">
					<img v-for="photo in profile" :src="'/photos/image-' + photo.photoId +'.png'" @click="deletePhoto(photo.photoId)" class="profile-image" :id="photo.photoId" v-bind:key="photo.photoId">
				</div>
			</div>
			<button class="btn btn-outline-secondary" id="delete-button" type="button" @click="toggleDeleteMode" disabled>Delete Photo</button>	
		</div>
		<div style="border: solid; border-color: grey; border-radius: 10px; padding: 20px; margin-right: 40vw; margin-top: 20px; margin-bottom: 20px; background-color: rgba(85,165,214, 0.4);">
			<p class="h3">Check your friends' profile</p>
			<div style="display: flex; flex-direction: row;">
				<input type="text" id="others-profile-input" style="margin-right: 10px;" class="form-control" placeholder="Whose profile do you want to visit?">
				<button class="btn btn-outline-secondary" type="button" @click="getUserProfile('others')">Visit</button>
			</div>	
		</div>
		<div id="image-grid" class="image-grid" v-if="this.toVisitProfile.length && !this.loading">
				<img v-for="photo in toVisitProfile" :src="'/photos/image-' + photo.photoId +'.png'" @click="deletePhoto(photo.photoId)" class="profile-image" :id="photo.photoId" v-bind:key="photo.photoId">
		</div>
		<p v-if="!this.toVisitProfile.length && this.profileRetrieved" class="h4" style="margin-left: 20px;">Oops...{{ profileRetrieved }} has not posted anything yet</p>
	</div>
</template>

<style>

.image-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    grid-gap: 3vw;
    padding: 10px;
}

.image-grid img {
    object-fit: contain;
    object-position: center;
}

.profile-image {
    background-color: transparent;
	border-radius: 10px;
	height: 40vh;
	width: 90%;
	height: auto;
	
}	

</style>



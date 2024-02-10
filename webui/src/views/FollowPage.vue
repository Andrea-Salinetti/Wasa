<script>
export default {
	data: function() {
		return {
			errormsg: null,
			positiveBanner: null,
			loading: false,
			some_data: null,
			userId: sessionStorage.getItem('userId'),
			username: sessionStorage.getItem('username'),
			logged: sessionStorage.getItem('logged'),
			profileRetrieved: false,
			profile: [],
			toVisitProfile: []
		}
	},
	methods: {
		async followUser() {
            this.loading = true;
            this.errormsg = null;
			const toFollowUsername = document.getElementById("follow-username-input").value;
			document.getElementById("follow-username-input").value = ""
			console.log(this.username)
            try {
                let response = await this.$axios.post("/users/"+sessionStorage.getItem('userId')+"/followed/?username="+ toFollowUsername, {}, {headers:{'Authorization': 'Bearer ' + sessionStorage.getItem('userId')}});
				if (response.status<400){
                console.log("User succesfully followed", response.data);
				this.positiveBanner = 'You are now following "' + toFollowUsername + '"'
				setTimeout(()=>{
					this.positiveBanner = null
				}, 3000)
				}
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },

		async banUser() {
            this.loading = true;
            this.errormsg = null;
			const toBanUsername = document.getElementById("ban-username-input").value;
			document.getElementById("ban-username-input").value = ""
            try {
                let response = await this.$axios.post("/users/"+sessionStorage.getItem('userId')+"/banned/?username="+ toBanUsername, {}, {headers:{'Authorization': 'Bearer ' + sessionStorage.getItem('userId')}});
				if (response.status<400){
                console.log("User succesfully followed", response.data);
				this.positiveBanner = '"' + toBanUsername + '" Has been banned'
				setTimeout(()=>{
					this.positiveBanner = null
				}, 3000)
				}
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },

		async unbanUser() {
            this.loading = true;
            this.errormsg = null;
			const toUnbanUsername = document.getElementById("unban-username-input").value;
			document.getElementById("unban-username-input").value = ""
            try {
				console.log("/users/"+sessionStorage.getItem('userId')+"/banned/"+toUnbanUsername)
                let response = await this.$axios.delete("/users/"+sessionStorage.getItem('userId')+"/banned/"+toUnbanUsername, {headers:{'Authorization': 'Bearer ' + sessionStorage.getItem('userId')}});
				if (response.status<400){
                console.log("User succesfully unbanned", response.data);
				this.positiveBanner = '"' + toUnbanUsername + '" Is not banned anymore'
				setTimeout(()=>{
					this.positiveBanner = null
				}, 3000)
				}
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },


		async unfollowUser() {
            this.loading = true;
            this.errormsg = null;
			const toUnfollowUsername = document.getElementById("unfollow-username-input").value;
			document.getElementById("unfollow-username-input").value = ""
            try {
				console.log("/users/"+sessionStorage.getItem('userId')+"/followed/"+toUnfollowUsername)
                let response = await this.$axios.delete("/users/"+sessionStorage.getItem('userId')+"/followed/"+toUnfollowUsername, {headers:{'Authorization': 'Bearer ' + sessionStorage.getItem('userId')}});
				if (response.status<400){
                console.log("User succesfully unfollowed", response.data);
				this.positiveBanner = 'You are not following "' + toUnfollowUsername + '" anymore'
				setTimeout(()=>{
					this.positiveBanner = null
				}, 3000)
				}
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },


		async getUserProfile(username){
			if (!sessionStorage.getItem('logged')){
				return
			}
			this.loading = true;
			this.errormsg = null;
			this.profileRetrieved = false;
			let toVisitUsername;

			toVisitUsername = document.getElementById('others-profile-input').value

			console.log(username)
			const config = {
				headers: {
				'Authorization': 'Bearer ' + sessionStorage.getItem('userId')
				}
			};
			try {
				let response = await this.$axios.get("/profiles/"+ toVisitUsername + "?userId=" + sessionStorage.getItem('userId'), config);

					this.toVisitProfile = response.data
					this.profileRetrieved = toVisitUsername
				console.log(this.profile)

			} catch (e) {
				this.errormsg = e.toString();
			}
			if (this.profile.length){
				const deleteButton = document.getElementById("delete-button")
				deleteButton.disabled = false
			}
			this.loading = false;
		},
		async likePhoto(photoId, i){
		try {
			let response = await this.$axios.post("/likes/?photoId="+ photoId +"&userId=" + sessionStorage.getItem('userId'),{},{headers:{
				"Authorization": "Bearer " + sessionStorage.getItem('userId')
			}});
			if (response.status<400) {
				console.log("photo liked")
				const button = document.getElementById('Lo'+photoId)
				button.id = response.data
				button.className= 'like-button-active'
				this.toVisitProfile[i].likeId = response.data
				this.toVisitProfile[i].likes++;
			}

		} catch (e) {
			this.errormsg = e.toString();
			console.log(e)
		}
	},

	async unlikePhoto(photoId, id, i){
		try {
			let response = await this.$axios.delete("/likes/" + id + "?userId=" + sessionStorage.getItem('userId'), {
				headers: {
					"Authorization": "Bearer " + sessionStorage.getItem('userId')
				}
			});

			if (response.status<400) {
				console.log("Photo disliked");
				const button = document.getElementById(id)
				button.id = 'Lo' + photoId
				button.className = 'like-button'
				this.toVisitProfile[i].likes--;
			}

		} catch (e) {
			this.errormsg = e.toString();
			console.error(e);
		}
	},

	async likeEventHandler(photoId, event, i) {
		this.loading = true;
		this.errormsg = null;
		const id = event.target.id
		console.log(event.target.id);

		if (id == 'Lo'+photoId) {
			this.likePhoto(photoId, i)
		} else {
			this.unlikePhoto(photoId, id, i)
		}
		},
	async commentPhoto(photoId, i) {
				this.loading = true;
				this.errormsg = null;
				const commentInput = document.getElementById('CI'+ photoId)
				const comment = commentInput.value
				if (!comment){
					return
				}
				const formData = new FormData();
				formData.append("content", comment)
				console.log(comment)
				commentInput.value = ""
				const config = {headers:{
						"Authorization": "Bearer " + sessionStorage.getItem('userId'),
						"Content-Type": "multipart/form-data"
					}}
				const url = "/comments/?photoId="+ photoId +"&userId=" + sessionStorage.getItem('userId')
				try {
					let response = await this.$axios.post(url, formData, config);
					if (response.status == 201) {
						console.log(response.data)
					}
				this.toVisitProfile[i].comments.push({
					content: comment,
					commentId: response.data,
					username: this.username
				})
				} catch (e) {
					this.errormsg = e.toString();
					console.log(e)
				}
				this.loading = false;
	},
	async uncommentPhoto(commentId){
	try {
			let response = await this.$axios.delete("/comments/" + commentId + "?userId=" + sessionStorage.getItem('userId'), {
				headers: {
					"Authorization": "Bearer " + sessionStorage.getItem('userId')
				}
			});

			if (true) {
				console.log("comment deleted");
				const content = document.getElementById('C'+commentId)
				const div = document.getElementById('D'+commentId)
				div.remove()
				content.remove()
			}

		} catch (e) {
			console.log("/likes/?photoId="+ photoId +"&userId=" + sessionStorage.getItem('userId'))
			this.errormsg = e.toString();
			console.error(e);
		}
	}
	},
	mounted() {
		//this.refresh()
	}
}
</script>

<template>
	<p v-if="!this.logged" class="positive-banner" style="margin-top: 40px;">In this page you will be able to manage your network, but please log-in first in the home page</p>
	<div v-if="this.logged">
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Your Network</h1>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<p v-if="positiveBanner" class="positive-banner">{{ this.positiveBanner }}</p>
		<p class="h3">Want to make new friends?</p>
		<div style="border: solid; border-color: grey; border-radius: 10px; padding: 20px; margin-bottom: 20px; margin-right: 40vw; background-color: rgba(208, 219, 252,1);" >
			<div style="display: flex; flex-direction: row;">
				<input type="text" id="follow-username-input" style="margin-right: 10px;" class="form-control" placeholder="Who do you want to follow?">
        		<button class="btn btn-outline-secondary" type="button" @click="followUser">Follow</button>
			</div>
			<div style="display: flex; flex-direction: row;">
				<input type="text" id="unfollow-username-input" style="margin-top: 20px; margin-right: 10px;" class="form-control" placeholder="Who do you want to unfollow?">
        		<button class="btn btn-outline-secondary" type="button" @click="unfollowUser" style="margin-top: 20px;">Unfollow</button>
			</div>
		</div>
		<p class="h3">Someone is bothering you?</p>
		<div style="border: solid; border-color: grey; border-radius: 10px; padding: 20px; margin-right: 40vw; background-color: rgba(100, 210, 180, 0.6);">
			<div style="display: flex; flex-direction: row;">
				<input type="text" id="ban-username-input" style="margin-right: 10px;" class="form-control" placeholder="Who do you want to ban?">
        		<button class="btn btn-outline-secondary" type="button" @click="banUser">Ban</button>
			</div>	
				<div style="display: flex; flex-direction: row;">
				<input type="text" id="unban-username-input" style="margin-top: 20px; margin-right: 10px;" class="form-control" placeholder="Who do you want to unban?">
        	<button class="btn btn-outline-secondary" type="button" style=" margin-top: 20px;" @click="unbanUser">Unban</button>
			</div>
		</div>
		<p class="h3" style="margin-top: 20px;">Check your friends' profile</p>
		<div style="border: solid; border-color: grey; border-radius: 10px; padding: 20px; margin-right: 40vw; margin-top: 10px; margin-bottom: 20px; background-color: rgba(85,165,214, 0.4);">
			<div style="display: flex; flex-direction: row;">
				<input type="text" id="others-profile-input" style="margin-right: 10px;" class="form-control" placeholder="Whose profile do you want to visit?">
				<button class="btn btn-outline-secondary" type="button" @click="getUserProfile('others')">Visit</button>
			</div>	
		</div>
		<div v-for="(photo, i) in toVisitProfile" class="post-container" v-bind:key="'for1'+photo.photoId">
					<div class="image-container">
						<p class="username-title"> {{ photo.username }}</p>
						<img :src="'/photos/image-' + photo.photoId + '.png'" class="profile-image" :id="photo.photoId"> 
					</div>
					<div class="buttons-container">
						<button :class="photo.likeId === '' ? 'like-button' : 'like-button-active'" :id="photo.likeId !== '' ? photo.likeId : 'Lo'+photo.photoId " @click="likeEventHandler(photo.photoId, $event, i)">{{photo.likes, i}} <svg class="feather" @click.stop><use href="/feather-sprite-v4.29.0.svg#heart"/></svg></button>
						<div>
							<input type="text" placeholder="Leave a Comment..." class="comment-input" :id="'CI'+photo.photoId">
							<button class="comment-button" :id="'C'+photo.photoId" @click="commentPhoto(photo.photoId, i)"><svg class="feather"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg> Comment</button>
						</div>
					</div>
					<div v-for="comment in photo.comments" class="comment-box" v-bind:key="'for2' + comment.commentId">
						<div style="flex-direction: row; display: flex; justify-content: space-between;" :id="'D'+comment.commentId">
							<p v-if="comment.username == this.username" class="comment-title">You</p>
							<p v-if="comment.username != this.username" class="comment-title">{{ comment.username }}</p>
							<button v-if="comment.username == this.username" :id="comment.commentId" class="delete-comment-button" @click="uncommentPhoto(comment.commentId)"><svg class="feather"><use href="/feather-sprite-v4.29.0.svg#x"/></svg></button>
			</div>
			<p class="comment-content" :id="'C'+comment.commentId">{{ comment.content }}</p>
		</div>
	  </div>
	  <p v-if="!this.toVisitProfile.length && this.profileRetrieved" class="h4" style="margin-left: 20px;">Oops...{{ profileRetrieved }} has not posted anything yet</p>
	</div>
</template>

<style>
</style>


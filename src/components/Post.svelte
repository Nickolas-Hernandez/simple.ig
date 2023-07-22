<script>
    import profileImage from '$lib/assets/profile-image.jpg'

    const postData = {
        user: "LuckyCharms",
        profileImage: profileImage
    }
</script>

<section>
    <div class="image-frame" style="background-image: url({postData.profileImage})">
        
    </div>
</section>
 

 <style>
    .image-frame {
        width: 100%;
        height: 400px;
        background-repeat: no-repeat;
        background-position: center;
        background-size: cover;
    }
 </style>
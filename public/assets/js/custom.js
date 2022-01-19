/* NAVIGATION AREA*/ 
$(function() {
    //navigtion
    $(window).on("scroll", function() {
        if($(window).scrollTop() > 40) {
            $("nav").addClass("navbar-active");
        } else {
           $("nav").removeClass("navbar-active");
        }
    });

    if($(window).scrollTop() > 40) {
        $("nav").addClass("navbar-active");
    }


    $('.navbar-collapse a, a.navbar-brand').click(function(){
        $(".navbar-collapse").collapse('hide');
    });
});


$('#mainSliderArea').owlCarousel({
    loop:true,
    margin:10,
    nav:false,
    dots:false,
    autoplay:true,
    autoplayTimeout:3000,
    autoplayHoverPause:true,
    responsive:{
        0:{
            items:1
        },
        600:{
            items:1
        },
        1000:{
            items:1
        }
    }
})



$('#dfctipper-testimonials').owlCarousel({
    loop:true,
    margin:10,
    nav:false,
    dots:false,
    autoplay:true,
    autoplayTimeout:5000,
    autoplayHoverPause:true,
    responsive:{
        0:{
            items:1
        },
        600:{
            items:1
        },
        1000:{
            items:1
        }
    }
})

$('body').on('click', '#closebox2', function(){
    $('.white_content2').hide();
    $('.black_overlay2').hide();
})


/*COUNTDOWN TIMER*/ 
function getTimeRemaining(endtime) {
    const total = Date.parse(endtime) - Date.parse(new Date());
    const seconds = Math.floor((total / 1000) % 60);
    const minutes = Math.floor((total / 1000 / 60) % 60);
    const hours = Math.floor((total / (1000 * 60 * 60)) % 24);
    const days = Math.floor(total / (1000 * 60 * 60 * 24));
    
    return {
      total,
      days,
      hours,
      minutes,
      seconds
    };
}
  
function initializeClock(id, endtime) {
    const clock = document.getElementById(id);
    const daysSpan = clock.querySelector('.days');
    const hoursSpan = clock.querySelector('.hours');
    const minutesSpan = clock.querySelector('.minutes');
    const secondsSpan = clock.querySelector('.seconds');
  
    function updateClock() {
      const t = getTimeRemaining(endtime);
  
      
      hoursSpan.innerHTML = ('0' + t.hours).slice(-2);
      minutesSpan.innerHTML = ('0' + t.minutes).slice(-2);
      secondsSpan.innerHTML = ('0' + t.seconds).slice(-2);
  
      if (t.total <= 0) {
        clearInterval(timeinterval);
      }
    }
  
    updateClock();
    const timeinterval = setInterval(updateClock, 1000);
}
  
const timeInHours = 12;
const currentTime = Date.parse(new Date());
const deadline = new Date(currentTime + timeInHours*60*60*1000);

// const deadline = new Date(Date.parse(new Date()) + 15 * 24 * 60 * 60 * 1000);
if($('#clockdiv').length > 0){
 initializeClock('clockdiv', deadline);   
}



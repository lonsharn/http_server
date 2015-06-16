"use strict";

(function () {
	var scr, ctx, pointer, drag, boxes, dropc = 0,
		objects   = [],
		contacts  = [],
        player = null,
		
	// ==== init script ====
	var init = function () {
		// ---- screen ----
		scr = new ge1doot.Screen({
			container: "screen",
			resize: function () {
				resize();
			}
		});
		ctx = scr.ctx;
		scr.resize();
		boxes = document.getElementById("textures").getElementsByTagName('img');
		// ---- pointer events ----
		pointer = new ge1doot.Pointer({
			down: function() {
				newBox();
			},
			up: function () {
				for (var i = 0, rb; rb = objects[i++];) rb.drag = false;
				drag = false;
			}
		});
		numIterations = 10;
		kTimeStep = 1/60;
		kGravity = 25;
		kFriction = 0.3;
		run();
	}

    var draw = function(){
    }

    var action = function(){
    }

    var calculation = function(){
    }

    var clear() = function(){
        ctx.clearRect(0, 0, scr.width, scr.height);
    }

	// ======== main loop ========
	var run = function () {
		clear();
        //碰撞
        action();
        calculation()
		draw();
		// ---- animation loop ----
		requestAnimFrame(run);
	}
	return {
		// ---- onload event ----
		load : function () {
			window.addEventListener('load', function () {
				init();
			}, false);
		}
	}
})().load();

using System.Collections;
using System.Collections.Generic;
using UnityEngine;

//class for fade-in and fade-out yellow color, used by highlighted muscules
//Fixed Update is used for static 60 runs per second. 
//Default glow delta is 0.006 which gives the best glow we could have achieved

public class glowColorFader : MonoBehaviour {

    public Color currentFadeColor;
    private bool rising;
    private float delta = 0.006f;
    void setCurrentEmissionColor()
    {
        if (rising)
        {
            float r = currentFadeColor.r;
            float g = currentFadeColor.g;
            currentFadeColor = new Color(r + delta, g + delta, 0);
            if (r > 0.7f) { rising = false; }
        }
        else
        {
            float r = currentFadeColor.r;
            float g = currentFadeColor.g;
            currentFadeColor = new Color(r - delta, g - delta, 0);
            if (r < 0.3f) { rising = true; }
        }
    }
    void FixedUpdate()//synchronized emission color is made in this loop
    {
        setCurrentEmissionColor();
    }
}

using System.Collections;
using System.Collections.Generic;
using UnityEngine;

//Class for storing data of targetable muscules, determining if they are selectable or not
//and handling bloom glow effect via using global fading Color from raycasted_muscle_behaviour.cs
public class targetable_muscle_properties : MonoBehaviour {

    public int musculeIndex;
    public bool targeted;
    public GameObject controller;

   

    void Update()
    {
        if (targeted)
        {
          foreach(MeshRenderer i in this.GetComponentsInChildren<MeshRenderer>())
            {
                i.material.SetColor("_EmissionColor", controller.GetComponent<glowColorFader>().currentFadeColor);
            }
        }
        else
        {
            foreach (MeshRenderer i in this.GetComponentsInChildren<MeshRenderer>())
            {
                i.material.SetColor("_EmissionColor", Color.black);
            }
        }
    }
}

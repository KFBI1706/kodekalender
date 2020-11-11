
fn main() {
    let ct = "tMlsioaplnKlflgiruKanliaebeLlkslikkpnerikTasatamkDpsdakeraBeIdaegptnuaKtmteorpuTaTtbtsesOHXxonibmksekaaoaKtrssegnveinRedlkkkroeekVtkekymmlooLnanoKtlstoepHrpeutdynfSneloietbol";
    let pt =  "HestTubaTrompetKattungeIpadBakeredskapDatamaskinTrekkspillLekebilKanarifuglKnallpistolMobiltelefonSydenturHoppeslottKanonLommelyktVekkerklokkeRedningsvestKaraokemaskinXboxOst";

    let mut chars : Vec<char> = ct.chars().collect();
    let len = chars.len();

    let mut out : Vec<char> = vec!['c'; len];

    for i in (3..len+1).step_by(3).rev() {
        for (j,e) in chars[i-3..i].iter().enumerate() {
            out[j+(174-i)] = *e;
        }
    }

    chars = out.to_vec();

    for i in (1..len).step_by(2) {
        out[i] = chars[i-1];
        out[i-1] = chars[i];
    }

    chars = out.to_vec();

    for i in 0..len {
        out[(i+len/2)%len] = chars[i];
    }

    let ans: String = out.into_iter().collect();

    assert_eq!(ans,pt);
    println!("{}",ans);

}

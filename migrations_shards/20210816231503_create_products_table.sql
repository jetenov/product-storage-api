-- +goose Up
-- +goose StatementBegin
create table products
(
    id                        text        not null primary key,

    commercial_category_id    int         not null,
    description_category_id   int         not null,

    name                      text,
    description               text,
    state                     text,
    barcode                   text,
    vat                       int,
    price                     int,

    images                    jsonb,
    attributes                jsonb,
    raw_data                  jsonb,
    meta                      jsonb,

    depth                     int,
    weight                    int,
    height                    int,
    width                     int,

    block_reason              int,
    skus                      int[],
    fulfillment               smallint,
    competitor_count          int,
    canonical_ids             text[],

    created_at                timestamptz not null default now(),
    updated_at                timestamptz not null default now()
)partition by hash (id);

comment on column products.id is 'Canonical id filed from PriceMonitoring';

create table products_part_0 partition of products for values with (modulus 256, remainder 0);
create table products_part_1 partition of products for values with (modulus 256, remainder 1);
create table products_part_2 partition of products for values with (modulus 256, remainder 2);
create table products_part_3 partition of products for values with (modulus 256, remainder 3);
create table products_part_4 partition of products for values with (modulus 256, remainder 4);
create table products_part_5 partition of products for values with (modulus 256, remainder 5);
create table products_part_6 partition of products for values with (modulus 256, remainder 6);
create table products_part_7 partition of products for values with (modulus 256, remainder 7);
create table products_part_8 partition of products for values with (modulus 256, remainder 8);
create table products_part_9 partition of products for values with (modulus 256, remainder 9);
create table products_part_10 partition of products for values with (modulus 256, remainder 10);
create table products_part_11 partition of products for values with (modulus 256, remainder 11);
create table products_part_12 partition of products for values with (modulus 256, remainder 12);
create table products_part_13 partition of products for values with (modulus 256, remainder 13);
create table products_part_14 partition of products for values with (modulus 256, remainder 14);
create table products_part_15 partition of products for values with (modulus 256, remainder 15);
create table products_part_16 partition of products for values with (modulus 256, remainder 16);
create table products_part_17 partition of products for values with (modulus 256, remainder 17);
create table products_part_18 partition of products for values with (modulus 256, remainder 18);
create table products_part_19 partition of products for values with (modulus 256, remainder 19);
create table products_part_20 partition of products for values with (modulus 256, remainder 20);
create table products_part_21 partition of products for values with (modulus 256, remainder 21);
create table products_part_22 partition of products for values with (modulus 256, remainder 22);
create table products_part_23 partition of products for values with (modulus 256, remainder 23);
create table products_part_24 partition of products for values with (modulus 256, remainder 24);
create table products_part_25 partition of products for values with (modulus 256, remainder 25);
create table products_part_26 partition of products for values with (modulus 256, remainder 26);
create table products_part_27 partition of products for values with (modulus 256, remainder 27);
create table products_part_28 partition of products for values with (modulus 256, remainder 28);
create table products_part_29 partition of products for values with (modulus 256, remainder 29);
create table products_part_30 partition of products for values with (modulus 256, remainder 30);
create table products_part_31 partition of products for values with (modulus 256, remainder 31);
create table products_part_32 partition of products for values with (modulus 256, remainder 32);
create table products_part_33 partition of products for values with (modulus 256, remainder 33);
create table products_part_34 partition of products for values with (modulus 256, remainder 34);
create table products_part_35 partition of products for values with (modulus 256, remainder 35);
create table products_part_36 partition of products for values with (modulus 256, remainder 36);
create table products_part_37 partition of products for values with (modulus 256, remainder 37);
create table products_part_38 partition of products for values with (modulus 256, remainder 38);
create table products_part_39 partition of products for values with (modulus 256, remainder 39);
create table products_part_40 partition of products for values with (modulus 256, remainder 40);
create table products_part_41 partition of products for values with (modulus 256, remainder 41);
create table products_part_42 partition of products for values with (modulus 256, remainder 42);
create table products_part_43 partition of products for values with (modulus 256, remainder 43);
create table products_part_44 partition of products for values with (modulus 256, remainder 44);
create table products_part_45 partition of products for values with (modulus 256, remainder 45);
create table products_part_46 partition of products for values with (modulus 256, remainder 46);
create table products_part_47 partition of products for values with (modulus 256, remainder 47);
create table products_part_48 partition of products for values with (modulus 256, remainder 48);
create table products_part_49 partition of products for values with (modulus 256, remainder 49);
create table products_part_50 partition of products for values with (modulus 256, remainder 50);
create table products_part_51 partition of products for values with (modulus 256, remainder 51);
create table products_part_52 partition of products for values with (modulus 256, remainder 52);
create table products_part_53 partition of products for values with (modulus 256, remainder 53);
create table products_part_54 partition of products for values with (modulus 256, remainder 54);
create table products_part_55 partition of products for values with (modulus 256, remainder 55);
create table products_part_56 partition of products for values with (modulus 256, remainder 56);
create table products_part_57 partition of products for values with (modulus 256, remainder 57);
create table products_part_58 partition of products for values with (modulus 256, remainder 58);
create table products_part_59 partition of products for values with (modulus 256, remainder 59);
create table products_part_60 partition of products for values with (modulus 256, remainder 60);
create table products_part_61 partition of products for values with (modulus 256, remainder 61);
create table products_part_62 partition of products for values with (modulus 256, remainder 62);
create table products_part_63 partition of products for values with (modulus 256, remainder 63);
create table products_part_64 partition of products for values with (modulus 256, remainder 64);
create table products_part_65 partition of products for values with (modulus 256, remainder 65);
create table products_part_66 partition of products for values with (modulus 256, remainder 66);
create table products_part_67 partition of products for values with (modulus 256, remainder 67);
create table products_part_68 partition of products for values with (modulus 256, remainder 68);
create table products_part_69 partition of products for values with (modulus 256, remainder 69);
create table products_part_70 partition of products for values with (modulus 256, remainder 70);
create table products_part_71 partition of products for values with (modulus 256, remainder 71);
create table products_part_72 partition of products for values with (modulus 256, remainder 72);
create table products_part_73 partition of products for values with (modulus 256, remainder 73);
create table products_part_74 partition of products for values with (modulus 256, remainder 74);
create table products_part_75 partition of products for values with (modulus 256, remainder 75);
create table products_part_76 partition of products for values with (modulus 256, remainder 76);
create table products_part_77 partition of products for values with (modulus 256, remainder 77);
create table products_part_78 partition of products for values with (modulus 256, remainder 78);
create table products_part_79 partition of products for values with (modulus 256, remainder 79);
create table products_part_80 partition of products for values with (modulus 256, remainder 80);
create table products_part_81 partition of products for values with (modulus 256, remainder 81);
create table products_part_82 partition of products for values with (modulus 256, remainder 82);
create table products_part_83 partition of products for values with (modulus 256, remainder 83);
create table products_part_84 partition of products for values with (modulus 256, remainder 84);
create table products_part_85 partition of products for values with (modulus 256, remainder 85);
create table products_part_86 partition of products for values with (modulus 256, remainder 86);
create table products_part_87 partition of products for values with (modulus 256, remainder 87);
create table products_part_88 partition of products for values with (modulus 256, remainder 88);
create table products_part_89 partition of products for values with (modulus 256, remainder 89);
create table products_part_90 partition of products for values with (modulus 256, remainder 90);
create table products_part_91 partition of products for values with (modulus 256, remainder 91);
create table products_part_92 partition of products for values with (modulus 256, remainder 92);
create table products_part_93 partition of products for values with (modulus 256, remainder 93);
create table products_part_94 partition of products for values with (modulus 256, remainder 94);
create table products_part_95 partition of products for values with (modulus 256, remainder 95);
create table products_part_96 partition of products for values with (modulus 256, remainder 96);
create table products_part_97 partition of products for values with (modulus 256, remainder 97);
create table products_part_98 partition of products for values with (modulus 256, remainder 98);
create table products_part_99 partition of products for values with (modulus 256, remainder 99);
create table products_part_100 partition of products for values with (modulus 256, remainder 100);
create table products_part_101 partition of products for values with (modulus 256, remainder 101);
create table products_part_102 partition of products for values with (modulus 256, remainder 102);
create table products_part_103 partition of products for values with (modulus 256, remainder 103);
create table products_part_104 partition of products for values with (modulus 256, remainder 104);
create table products_part_105 partition of products for values with (modulus 256, remainder 105);
create table products_part_106 partition of products for values with (modulus 256, remainder 106);
create table products_part_107 partition of products for values with (modulus 256, remainder 107);
create table products_part_108 partition of products for values with (modulus 256, remainder 108);
create table products_part_109 partition of products for values with (modulus 256, remainder 109);
create table products_part_110 partition of products for values with (modulus 256, remainder 110);
create table products_part_111 partition of products for values with (modulus 256, remainder 111);
create table products_part_112 partition of products for values with (modulus 256, remainder 112);
create table products_part_113 partition of products for values with (modulus 256, remainder 113);
create table products_part_114 partition of products for values with (modulus 256, remainder 114);
create table products_part_115 partition of products for values with (modulus 256, remainder 115);
create table products_part_116 partition of products for values with (modulus 256, remainder 116);
create table products_part_117 partition of products for values with (modulus 256, remainder 117);
create table products_part_118 partition of products for values with (modulus 256, remainder 118);
create table products_part_119 partition of products for values with (modulus 256, remainder 119);
create table products_part_120 partition of products for values with (modulus 256, remainder 120);
create table products_part_121 partition of products for values with (modulus 256, remainder 121);
create table products_part_122 partition of products for values with (modulus 256, remainder 122);
create table products_part_123 partition of products for values with (modulus 256, remainder 123);
create table products_part_124 partition of products for values with (modulus 256, remainder 124);
create table products_part_125 partition of products for values with (modulus 256, remainder 125);
create table products_part_126 partition of products for values with (modulus 256, remainder 126);
create table products_part_127 partition of products for values with (modulus 256, remainder 127);
create table products_part_128 partition of products for values with (modulus 256, remainder 128);
create table products_part_129 partition of products for values with (modulus 256, remainder 129);
create table products_part_130 partition of products for values with (modulus 256, remainder 130);
create table products_part_131 partition of products for values with (modulus 256, remainder 131);
create table products_part_132 partition of products for values with (modulus 256, remainder 132);
create table products_part_133 partition of products for values with (modulus 256, remainder 133);
create table products_part_134 partition of products for values with (modulus 256, remainder 134);
create table products_part_135 partition of products for values with (modulus 256, remainder 135);
create table products_part_136 partition of products for values with (modulus 256, remainder 136);
create table products_part_137 partition of products for values with (modulus 256, remainder 137);
create table products_part_138 partition of products for values with (modulus 256, remainder 138);
create table products_part_139 partition of products for values with (modulus 256, remainder 139);
create table products_part_140 partition of products for values with (modulus 256, remainder 140);
create table products_part_141 partition of products for values with (modulus 256, remainder 141);
create table products_part_142 partition of products for values with (modulus 256, remainder 142);
create table products_part_143 partition of products for values with (modulus 256, remainder 143);
create table products_part_144 partition of products for values with (modulus 256, remainder 144);
create table products_part_145 partition of products for values with (modulus 256, remainder 145);
create table products_part_146 partition of products for values with (modulus 256, remainder 146);
create table products_part_147 partition of products for values with (modulus 256, remainder 147);
create table products_part_148 partition of products for values with (modulus 256, remainder 148);
create table products_part_149 partition of products for values with (modulus 256, remainder 149);
create table products_part_150 partition of products for values with (modulus 256, remainder 150);
create table products_part_151 partition of products for values with (modulus 256, remainder 151);
create table products_part_152 partition of products for values with (modulus 256, remainder 152);
create table products_part_153 partition of products for values with (modulus 256, remainder 153);
create table products_part_154 partition of products for values with (modulus 256, remainder 154);
create table products_part_155 partition of products for values with (modulus 256, remainder 155);
create table products_part_156 partition of products for values with (modulus 256, remainder 156);
create table products_part_157 partition of products for values with (modulus 256, remainder 157);
create table products_part_158 partition of products for values with (modulus 256, remainder 158);
create table products_part_159 partition of products for values with (modulus 256, remainder 159);
create table products_part_160 partition of products for values with (modulus 256, remainder 160);
create table products_part_161 partition of products for values with (modulus 256, remainder 161);
create table products_part_162 partition of products for values with (modulus 256, remainder 162);
create table products_part_163 partition of products for values with (modulus 256, remainder 163);
create table products_part_164 partition of products for values with (modulus 256, remainder 164);
create table products_part_165 partition of products for values with (modulus 256, remainder 165);
create table products_part_166 partition of products for values with (modulus 256, remainder 166);
create table products_part_167 partition of products for values with (modulus 256, remainder 167);
create table products_part_168 partition of products for values with (modulus 256, remainder 168);
create table products_part_169 partition of products for values with (modulus 256, remainder 169);
create table products_part_170 partition of products for values with (modulus 256, remainder 170);
create table products_part_171 partition of products for values with (modulus 256, remainder 171);
create table products_part_172 partition of products for values with (modulus 256, remainder 172);
create table products_part_173 partition of products for values with (modulus 256, remainder 173);
create table products_part_174 partition of products for values with (modulus 256, remainder 174);
create table products_part_175 partition of products for values with (modulus 256, remainder 175);
create table products_part_176 partition of products for values with (modulus 256, remainder 176);
create table products_part_177 partition of products for values with (modulus 256, remainder 177);
create table products_part_178 partition of products for values with (modulus 256, remainder 178);
create table products_part_179 partition of products for values with (modulus 256, remainder 179);
create table products_part_180 partition of products for values with (modulus 256, remainder 180);
create table products_part_181 partition of products for values with (modulus 256, remainder 181);
create table products_part_182 partition of products for values with (modulus 256, remainder 182);
create table products_part_183 partition of products for values with (modulus 256, remainder 183);
create table products_part_184 partition of products for values with (modulus 256, remainder 184);
create table products_part_185 partition of products for values with (modulus 256, remainder 185);
create table products_part_186 partition of products for values with (modulus 256, remainder 186);
create table products_part_187 partition of products for values with (modulus 256, remainder 187);
create table products_part_188 partition of products for values with (modulus 256, remainder 188);
create table products_part_189 partition of products for values with (modulus 256, remainder 189);
create table products_part_190 partition of products for values with (modulus 256, remainder 190);
create table products_part_191 partition of products for values with (modulus 256, remainder 191);
create table products_part_192 partition of products for values with (modulus 256, remainder 192);
create table products_part_193 partition of products for values with (modulus 256, remainder 193);
create table products_part_194 partition of products for values with (modulus 256, remainder 194);
create table products_part_195 partition of products for values with (modulus 256, remainder 195);
create table products_part_196 partition of products for values with (modulus 256, remainder 196);
create table products_part_197 partition of products for values with (modulus 256, remainder 197);
create table products_part_198 partition of products for values with (modulus 256, remainder 198);
create table products_part_199 partition of products for values with (modulus 256, remainder 199);
create table products_part_200 partition of products for values with (modulus 256, remainder 200);
create table products_part_201 partition of products for values with (modulus 256, remainder 201);
create table products_part_202 partition of products for values with (modulus 256, remainder 202);
create table products_part_203 partition of products for values with (modulus 256, remainder 203);
create table products_part_204 partition of products for values with (modulus 256, remainder 204);
create table products_part_205 partition of products for values with (modulus 256, remainder 205);
create table products_part_206 partition of products for values with (modulus 256, remainder 206);
create table products_part_207 partition of products for values with (modulus 256, remainder 207);
create table products_part_208 partition of products for values with (modulus 256, remainder 208);
create table products_part_209 partition of products for values with (modulus 256, remainder 209);
create table products_part_210 partition of products for values with (modulus 256, remainder 210);
create table products_part_211 partition of products for values with (modulus 256, remainder 211);
create table products_part_212 partition of products for values with (modulus 256, remainder 212);
create table products_part_213 partition of products for values with (modulus 256, remainder 213);
create table products_part_214 partition of products for values with (modulus 256, remainder 214);
create table products_part_215 partition of products for values with (modulus 256, remainder 215);
create table products_part_216 partition of products for values with (modulus 256, remainder 216);
create table products_part_217 partition of products for values with (modulus 256, remainder 217);
create table products_part_218 partition of products for values with (modulus 256, remainder 218);
create table products_part_219 partition of products for values with (modulus 256, remainder 219);
create table products_part_220 partition of products for values with (modulus 256, remainder 220);
create table products_part_221 partition of products for values with (modulus 256, remainder 221);
create table products_part_222 partition of products for values with (modulus 256, remainder 222);
create table products_part_223 partition of products for values with (modulus 256, remainder 223);
create table products_part_224 partition of products for values with (modulus 256, remainder 224);
create table products_part_225 partition of products for values with (modulus 256, remainder 225);
create table products_part_226 partition of products for values with (modulus 256, remainder 226);
create table products_part_227 partition of products for values with (modulus 256, remainder 227);
create table products_part_228 partition of products for values with (modulus 256, remainder 228);
create table products_part_229 partition of products for values with (modulus 256, remainder 229);
create table products_part_230 partition of products for values with (modulus 256, remainder 230);
create table products_part_231 partition of products for values with (modulus 256, remainder 231);
create table products_part_232 partition of products for values with (modulus 256, remainder 232);
create table products_part_233 partition of products for values with (modulus 256, remainder 233);
create table products_part_234 partition of products for values with (modulus 256, remainder 234);
create table products_part_235 partition of products for values with (modulus 256, remainder 235);
create table products_part_236 partition of products for values with (modulus 256, remainder 236);
create table products_part_237 partition of products for values with (modulus 256, remainder 237);
create table products_part_238 partition of products for values with (modulus 256, remainder 238);
create table products_part_239 partition of products for values with (modulus 256, remainder 239);
create table products_part_240 partition of products for values with (modulus 256, remainder 240);
create table products_part_241 partition of products for values with (modulus 256, remainder 241);
create table products_part_242 partition of products for values with (modulus 256, remainder 242);
create table products_part_243 partition of products for values with (modulus 256, remainder 243);
create table products_part_244 partition of products for values with (modulus 256, remainder 244);
create table products_part_245 partition of products for values with (modulus 256, remainder 245);
create table products_part_246 partition of products for values with (modulus 256, remainder 246);
create table products_part_247 partition of products for values with (modulus 256, remainder 247);
create table products_part_248 partition of products for values with (modulus 256, remainder 248);
create table products_part_249 partition of products for values with (modulus 256, remainder 249);
create table products_part_250 partition of products for values with (modulus 256, remainder 250);
create table products_part_251 partition of products for values with (modulus 256, remainder 251);
create table products_part_252 partition of products for values with (modulus 256, remainder 252);
create table products_part_253 partition of products for values with (modulus 256, remainder 253);
create table products_part_254 partition of products for values with (modulus 256, remainder 254);
create table products_part_255 partition of products for values with (modulus 256, remainder 255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table products;
-- +goose StatementEnd
